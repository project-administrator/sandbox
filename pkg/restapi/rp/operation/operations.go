/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package operation

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hyperledger/aries-framework-go/spi/storage"
	"github.com/trustbloc/edge-core/pkg/log"
	edgesvcops "github.com/trustbloc/edge-service/pkg/restapi/verifier/operation"

	"github.com/trustbloc/sandbox/pkg/internal/common/support"
	oidcclient "github.com/trustbloc/sandbox/pkg/restapi/internal/common/oidc"
)

const (
	httpContentTypeJSON = "application/json"

	// api paths
	verifyVPPath           = "/verifyPresentation"
	oauth2GetRequestPath   = "/oauth2/request"
	oauth2CallbackPath     = "/oauth2/callback"
	verifyPresentationPath = "/verify/presentation"
	verifyCredentialPath   = "/verify/credential"

	// api path params
	scopeQueryParam    = "scope"
	flowQueryParam     = "flow"
	demoTypeQueryParam = "demoType"

	// edge-service verifier endpoints
	verifyPresentationURLFormat = "/%s" + "/verifier/presentations/verify"

	// edge-service verifier endpoints
	verifyCredentialURLFormat = "/%s" + "/verifier/credentials/verify"

	// TODO https://github.com/trustbloc/sandbox/issues/352 Configure verifier profiles in Verifier page
	verifierProfileID = "trustbloc-verifier"

	vcsVerifierRequestTokenName = "vcs_verifier" //nolint: gosec

	transientStoreName = "rp-rest-transient"
	flowTypeCookie     = "flowType"
	waciDemoType       = "waci"
)

var logger = log.New("sandbox-rp-restapi")

// Handler http handler for each controller API endpoint
type Handler interface {
	Path() string
	Method() string
	Handle() http.HandlerFunc
}

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type oidcClient interface {
	CreateOIDCRequest(state, scope string) (string, error)
	HandleOIDCCallback(reqContext context.Context, code string) ([]byte, error)
}

// Operation defines handlers
type Operation struct {
	handlers       []Handler
	vpHTML         string
	didCommVpHTML  string
	vcsURL         string
	client         httpClient
	requestTokens  map[string]string
	transientStore storage.Store
	tlsConfig      *tls.Config
	oidcClient     oidcClient
	waciOIDCClient oidcClient
}

// Config defines configuration for rp operations
type Config struct {
	VPHTML                 string
	DIDCOMMVPHTML          string
	VCSURL                 string
	TLSConfig              *tls.Config
	RequestTokens          map[string]string
	OIDCProviderURL        string
	OIDCClientID           string
	OIDCClientSecret       string
	OIDCCallbackURL        string
	TransientStoreProvider storage.Provider
	WACIOIDCProviderURL    string
	WACIOIDCClientID       string
	WACIOIDCClientSecret   string
	WACIOIDCCallbackURL    string
}

// vc struct used to return vc data to html
type vc struct {
	Data     string `json:"data"`
	Msg      string `json:"msg"`
	FlowType string `json:"flowType"`
}

type createOIDCRequestResponse struct {
	Request  string `json:"request"`
	FlowType string `json:"flowType"`
}

// New returns rp operation instance
func New(config *Config) (*Operation, error) {
	svc := &Operation{
		vpHTML:        config.VPHTML,
		didCommVpHTML: config.DIDCOMMVPHTML,
		vcsURL:        config.VCSURL,
		client:        &http.Client{Transport: &http.Transport{TLSClientConfig: config.TLSConfig}},
		requestTokens: config.RequestTokens,
		tlsConfig:     config.TLSConfig,
	}

	var err error

	svc.oidcClient, err = oidcclient.New(&oidcclient.Config{
		OIDCClientID:     config.OIDCClientID,
		OIDCClientSecret: config.OIDCClientSecret, OIDCCallbackURL: config.OIDCCallbackURL,
		OIDCProviderURL: config.OIDCProviderURL, TLSConfig: config.TLSConfig,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create oidc client : %w", err)
	}

	svc.waciOIDCClient, err = oidcclient.New(&oidcclient.Config{
		OIDCClientID:     config.WACIOIDCClientID,
		OIDCClientSecret: config.WACIOIDCClientSecret, OIDCCallbackURL: config.WACIOIDCCallbackURL,
		OIDCProviderURL: config.WACIOIDCProviderURL, TLSConfig: config.TLSConfig,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create oidc client : %w", err)
	}

	svc.transientStore, err = createStore(config.TransientStoreProvider)
	if err != nil {
		return nil, fmt.Errorf("failed to create store : %w", err)
	}

	svc.registerHandler()

	return svc, nil
}

// registerHandler register handlers to be exposed from this service as REST API endpoints
func (c *Operation) registerHandler() {
	// Add more protocol endpoints here to expose them as controller API endpoints
	c.handlers = []Handler{
		support.NewHTTPHandler(verifyVPPath, http.MethodPost, c.verifyVP),
		support.NewHTTPHandler(oauth2GetRequestPath, http.MethodGet, c.createOIDCRequest),
		support.NewHTTPHandler(oauth2CallbackPath, http.MethodGet, c.handleOIDCCallback),

		support.NewHTTPHandler(verifyPresentationPath, http.MethodPost, c.verifyPresentation),
		support.NewHTTPHandler(verifyCredentialPath, http.MethodPost, c.verifyCredential),
	}
}

// GetRESTHandlers get all controller API handler available for this service
func (c *Operation) GetRESTHandlers() []Handler {
	return c.handlers
}

func (c *Operation) verifyPresentation(w http.ResponseWriter, r *http.Request) {
	req := &verifyPresentationRequest{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		c.writeErrorResponse(w, http.StatusBadRequest,
			fmt.Sprintf("failed to decode request: %s", err.Error()))

		return
	}

	vpReq := edgesvcops.VerifyPresentationRequest{
		Presentation: req.VP,
		Opts: &edgesvcops.VerifyPresentationOptions{
			Checks:    req.Checks,
			Challenge: req.Challenge,
			Domain:    req.Domain,
		},
	}

	resp, err := c.callVerifyPresentation(vpReq)
	if err != nil {
		c.writeErrorResponse(w, http.StatusBadRequest,
			fmt.Sprintf("failed to verify vp: %s", err.Error()))

		return
	}

	if resp.StatusCode != http.StatusOK {
		respBytes, respErr := ioutil.ReadAll(resp.Body)
		if respErr != nil {
			c.writeErrorResponse(w, http.StatusBadRequest,
				fmt.Sprintf("failed to read verify presentation resp : %s", err.Error()))

			return
		}

		defer func() {
			e := resp.Body.Close()
			if e != nil {
				logger.Errorf("closing response body failed: %v", e)
			}
		}()

		c.writeErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("failed to verify presentation: %s", string(respBytes)))

		return
	}

	c.writeResponse(w, http.StatusOK, []byte(""))
}

func (c *Operation) verifyCredential(w http.ResponseWriter, r *http.Request) {
	req := &verifyCredentialRequest{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		c.writeErrorResponse(w, http.StatusBadRequest,
			fmt.Sprintf("failed to decode request: %s", err.Error()))

		return
	}

	vcReq := edgesvcops.CredentialsVerificationRequest{
		Credential: req.VC,
		Opts: &edgesvcops.CredentialsVerificationOptions{
			Checks: req.Checks,
		},
	}

	resp, err := c.callVerifyCredential(vcReq)
	if err != nil {
		c.writeErrorResponse(w, http.StatusBadRequest,
			fmt.Sprintf("failed to verify vc: %s", err.Error()))

		return
	}

	if resp.StatusCode != http.StatusOK {
		respBytes, respErr := ioutil.ReadAll(resp.Body)
		if respErr != nil {
			c.writeErrorResponse(w, http.StatusBadRequest,
				fmt.Sprintf("failed to read verify credentail resp : %s", err.Error()))

			return
		}

		defer func() {
			e := resp.Body.Close()
			if e != nil {
				logger.Errorf("closing response body failed: %v", e)
			}
		}()

		c.writeErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("failed to verify credential: %s", string(respBytes)))

		return
	}

	c.writeResponse(w, http.StatusOK, []byte(""))
}

// verifyVP
func (c *Operation) verifyVP(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		c.writeErrorResponse(w, http.StatusInternalServerError,
			fmt.Sprintf("failed to parse form: %s", err.Error()))

		return
	}

	inputData := "vpDataInput"
	checks := strings.Split(r.Form.Get("checks"), ",")
	domain := r.Form.Get("domain")
	challenge := r.Form.Get("challenge")

	req := edgesvcops.VerifyPresentationRequest{
		Presentation: []byte(r.Form.Get(inputData)),
		Opts: &edgesvcops.VerifyPresentationOptions{
			Checks:    checks,
			Challenge: challenge,
			Domain:    domain,
		},
	}

	c.verify(req, inputData, c.vpHTML, w, r)
}

func (c *Operation) createOIDCRequest(w http.ResponseWriter, r *http.Request) { // nolint: funlen
	scope := r.URL.Query().Get(scopeQueryParam)
	if scope == "" {
		c.writeErrorResponse(w, http.StatusBadRequest, "missing scope")

		return
	}

	flowType := r.URL.Query().Get(flowQueryParam)
	if flowType == "" {
		c.writeErrorResponse(w, http.StatusBadRequest, "missing flow type")

		return
	}

	demoType := r.URL.Query().Get(demoTypeQueryParam)

	// TODO validate scope
	state := uuid.New().String()

	adapterOIDCClient := c.oidcClient

	if demoType == waciDemoType {
		adapterOIDCClient = c.waciOIDCClient
	}

	redirectURL, err := adapterOIDCClient.CreateOIDCRequest(state, scope)
	if err != nil {
		c.writeErrorResponse(w,
			http.StatusInternalServerError, fmt.Sprintf("failed to create oidc request : %s", err))

		return
	}

	response, err := json.Marshal(&createOIDCRequestResponse{
		Request:  redirectURL,
		FlowType: flowType,
	})
	if err != nil {
		c.writeErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("failed to marshal response : %s", err))

		return
	}

	err = c.transientStore.Put(state, []byte(demoType))
	if err != nil {
		c.writeErrorResponse(w,
			http.StatusInternalServerError, fmt.Sprintf("failed to write state to transient store : %s", err))

		return
	}

	expire := time.Now().AddDate(0, 0, 1)
	cookie := http.Cookie{Name: flowTypeCookie, Value: flowType, Expires: expire}
	http.SetCookie(w, &cookie)

	w.Header().Set("content-type", "application/json")

	_, err = w.Write(response)
	if err != nil {
		logger.Errorf("failed to write response : %s", err)
	}
}

func (c *Operation) handleOIDCCallback(w http.ResponseWriter, r *http.Request) {
	state := r.URL.Query().Get("state")
	if state == "" {
		logger.Errorf("missing state")
		c.didcommDemoResult(w, "missing state", "")

		return
	}

	code := r.URL.Query().Get("code")
	if code == "" {
		logger.Errorf("missing code")
		c.didcommDemoResult(w, "missing code", "")

		return
	}

	flowTypeCookie, err := r.Cookie(flowTypeCookie)
	if err != nil && !errors.Is(err, http.ErrNoCookie) {
		logger.Errorf(err.Error())
		c.writeErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("failed to get flowType cookie: %s",
			err.Error()))

		return
	}

	demoType, err := c.transientStore.Get(state)
	if errors.Is(err, storage.ErrDataNotFound) {
		logger.Errorf("invalid state parameter")
		c.didcommDemoResult(w, "invalid state parameter", "")

		return
	}

	if err != nil {
		logger.Errorf("failed to query transient store for state : %s", err)
		c.didcommDemoResult(w, fmt.Sprintf("failed to query transient store for state : %s", err), "")

		return
	}

	adapterOIDCClient := c.oidcClient
	if string(demoType) == waciDemoType {
		adapterOIDCClient = c.waciOIDCClient
	}

	data, err := adapterOIDCClient.HandleOIDCCallback(r.Context(), code)
	if err != nil {
		logger.Errorf("failed to handle oidc callback : %s", err)
		c.didcommDemoResult(w, fmt.Sprintf("failed to handle oidc callback: %s", err), "")

		return
	}

	c.didcommDemoResult(w, string(data), flowTypeCookie.Value)
}

func (c *Operation) didcommDemoResult(w http.ResponseWriter, data, flowType string) {
	t, err := template.ParseFiles(c.didCommVpHTML)
	if err != nil {
		c.writeErrorResponse(w, http.StatusInternalServerError,
			fmt.Sprintf("unable to load html: %s", err.Error()))

		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err := t.Execute(w, vc{Data: data, FlowType: flowType}); err != nil {
		logger.Errorf(fmt.Sprintf("failed execute html template: %s", err.Error()))
	}
}

// verify function verifies the input data and parse the response to provided template
func (c *Operation) verify(verifyReq interface{}, inputData, htmlTemplate string,
	w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(htmlTemplate)
	if err != nil {
		c.writeErrorResponse(w, http.StatusInternalServerError,
			fmt.Sprintf("unable to load html: %s", err.Error()))

		return
	}

	resp, httpErr := c.callVerifyPresentation(verifyReq)
	if httpErr != nil {
		c.writeErrorResponse(w, http.StatusBadRequest,
			fmt.Sprintf("failed to verify: %s", httpErr.Error()))

		if err := t.Execute(w, vc{Msg: "Oops verification is failed, Try again"}); err != nil {
			logger.Errorf(fmt.Sprintf("failed execute html template: %s", err.Error()))
		}

		return
	}

	if resp.StatusCode != http.StatusOK { //nolint:nestif
		var failedMsg string

		respBytes, respErr := ioutil.ReadAll(resp.Body)
		if respErr != nil {
			failedMsg = fmt.Sprintf("failed to read response body: %s", respErr)
		} else {
			failedMsg = string(respBytes)
			isStatusRevoked := checkVCStatus(failedMsg, w, t)
			if isStatusRevoked {
				return
			}
		}

		defer func() {
			e := resp.Body.Close()
			if e != nil {
				logger.Errorf("closing response body failed: %v", e)
			}
		}()

		c.writeErrorResponse(w, resp.StatusCode, fmt.Sprintf("failed to verify: %s", failedMsg))

		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err := t.Execute(w, vc{Msg: "Successfully verified", Data: r.Form.Get(inputData)}); err != nil {
		logger.Errorf(fmt.Sprintf("failed execute html template: %s", err.Error()))
	}
}

func (c *Operation) callVerifyPresentation(verifyReq interface{}) (*http.Response, error) {
	endpoint := fmt.Sprintf(verifyPresentationURLFormat, verifierProfileID)

	reqBytes, err := json.Marshal(verifyReq)
	if err != nil {
		return nil, fmt.Errorf("unmarshal request : %w", err)
	}

	return c.sendHTTPRequest(http.MethodPost, c.vcsURL+endpoint, reqBytes, httpContentTypeJSON,
		c.requestTokens[vcsVerifierRequestTokenName])
}

func (c *Operation) callVerifyCredential(verifyReq interface{}) (*http.Response, error) {
	endpoint := fmt.Sprintf(verifyCredentialURLFormat, verifierProfileID)

	reqBytes, err := json.Marshal(verifyReq)
	if err != nil {
		return nil, fmt.Errorf("unmarshal request : %w", err)
	}

	return c.sendHTTPRequest(http.MethodPost, c.vcsURL+endpoint, reqBytes, httpContentTypeJSON,
		c.requestTokens[vcsVerifierRequestTokenName])
}

func checkVCStatus(failedMsg string, rw io.Writer, t *template.Template) bool {
	isStatusRevoked := checkSubstrings(failedMsg, "Revoked")

	if isStatusRevoked {
		if err := t.Execute(rw, vc{Msg: "Oops verification is failed. VC is revoked"}); err != nil {
			logger.Errorf(fmt.Sprintf("failed execute html template: %s", err.Error()))
		}
	}

	return isStatusRevoked
}

func checkSubstrings(str string, subs ...string) bool {
	isCompleteMatch := false

	for _, sub := range subs {
		if strings.Contains(str, sub) {
			isCompleteMatch = true
		}
	}

	return isCompleteMatch
}

func (c *Operation) sendHTTPRequest(method, url string, body []byte, contentType,
	token string) (*http.Response, error) {
	logger.Infof("send http request : url=%s methdod=%s body=%s", url, method, string(body))

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}

	if token != "" {
		req.Header.Add("Authorization", "Bearer "+c.requestTokens[vcsVerifierRequestTokenName])
	}

	if contentType != "" {
		req.Header.Add("Content-Type", contentType)
	}

	return c.client.Do(req)
}

// writeResponse writes interface value to response
func (c *Operation) writeErrorResponse(rw http.ResponseWriter, status int, msg string) {
	logger.Errorf(msg)

	rw.WriteHeader(status)

	write := rw.Write
	if _, err := write([]byte(msg)); err != nil {
		logger.Errorf("Unable to send error message, %s", err)
	}
}

func createStore(p storage.Provider) (storage.Store, error) {
	return p.OpenStore(transientStoreName)
}

// writeResponse writes interface value to response
func (c *Operation) writeResponse(rw http.ResponseWriter, status int, data []byte) {
	rw.WriteHeader(status)

	if _, err := rw.Write(data); err != nil {
		logger.Errorf("Unable to send error message, %s", err)
	}
}
