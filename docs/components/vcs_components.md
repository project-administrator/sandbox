# TrustBloc VCS Components

## v0.1.7
### Architecture Diagram
![VCS Architecture diagram v0.1.7](../images/vcs_component_diagram_v0.1.7.svg)

TODO : Add link to DID Orb Component Diagram

### Component Details
#### Core
| Component                   | Source URL                                                                 | k8s Configuration                                                                                |
|-----------------------------|----------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------|
| Router                      | [Source](https://github.com/trustbloc/hub-router)                          | [k8s](https://github.com/trustbloc/k8s/tree/b7254443d1efdbb7f7819aeeffff687a38f40706/hub-router) |
| EDV                         | [Source](https://github.com/trustbloc/edv)                                 | [k8s](https://github.com/trustbloc/k8s/tree/b7254443d1efdbb7f7819aeeffff687a38f40706/edv)        |
| EDV Oathkeeper              | [Source](https://github.com/ory/oathkeeper)                                | [k8s](https://github.com/trustbloc/k8s/tree/b7254443d1efdbb7f7819aeeffff687a38f40706/edv)        |
| Ops KMS                     | [Source](https://github.com/trustbloc/kms)                                 | [k8s](https://github.com/trustbloc/k8s/tree/b7254443d1efdbb7f7819aeeffff687a38f40706/kms)        |
| Ops KMS Oathkeeper          | [Source](https://github.com/ory/oathkeeper )                               | [k8s](https://github.com/trustbloc/k8s/tree/b7254443d1efdbb7f7819aeeffff687a38f40706/kms)        |
| Authz KMS                   | [Source](https://github.com/trustbloc/kms)                                 | [k8s](https://github.com/trustbloc/k8s/tree/b7254443d1efdbb7f7819aeeffff687a38f40706/kms)        |
| Authz KMS Oathkeeper        | [Source](https://github.com/ory/oathkeeper )                               | [k8s](https://github.com/trustbloc/k8s/tree/b7254443d1efdbb7f7819aeeffff687a38f40706/kms)        |
| TrustBloc SignIn (Hub-Auth) | [Source](https://github.com/trustbloc/hub-auth)                            | [k8s](https://github.com/trustbloc/k8s/tree/b7254443d1efdbb7f7819aeeffff687a38f40706/hub-auth)   |
| TrustBloc SignIn (Hydra)    | [Source](https://github.com/ory/hydra)                                     | [k8s](https://github.com/trustbloc/k8s/tree/b7254443d1efdbb7f7819aeeffff687a38f40706/hub-auth)   |
| DID Resolver                | [Source](https://github.com/trustbloc/edge-service/tree/main/cmd/did-rest) | [k8s](https://github.com/trustbloc/k8s/tree/b7254443d1efdbb7f7819aeeffff687a38f40706/resolver)   |
| DID Orb                     | [Source](https://github.com/trustbloc/orb)                                 | [k8s](https://github.com/trustbloc/k8s/tree/b7254443d1efdbb7f7819aeeffff687a38f40706/orb)        |


#### Wallet
| Component     | Source URL                                    | k8s Configuration                                                                                   |
|---------------|-----------------------------------------------|-----------------------------------------------------------------------------------------------------|
| Wallet Web    | [Source](https://github.com/trustbloc/wallet) | [k8s](https://github.com/trustbloc/k8s/tree/b7254443d1efdbb7f7819aeeffff687a38f40706/wallet-web)    |
| Wallet Server | [Source](https://github.com/trustbloc/wallet) | [k8s](https://github.com/trustbloc/k8s/tree/b7254443d1efdbb7f7819aeeffff687a38f40706/wallet-server) |


#### Services
| Component           | Source URL                                          | k8s Configuration                                                                         |
|---------------------|-----------------------------------------------------|-------------------------------------------------------------------------------------------|
| Issuer VC Service   | [Source](https://github.com/trustbloc/edge-service) | [k8s](https://github.com/trustbloc/k8s/tree/b7254443d1efdbb7f7819aeeffff687a38f40706/vcs) |
| Verifier VC Service | [Source](https://github.com/trustbloc/edge-service) | [k8s](https://github.com/trustbloc/k8s/tree/b7254443d1efdbb7f7819aeeffff687a38f40706/vcs) |

#### Demo
| Component         | Source URL                                     | k8s Configuration                                                                                           |
|-------------------|------------------------------------------------|-------------------------------------------------------------------------------------------------------------|
| Issuer            | [Source](https://github.com/trustbloc/sandbox) | [k8s](https://github.com/trustbloc/sandbox/tree/59b6dee1552f7afdea3a8b6f804e3dfdfcc0837e/k8s/issuer)        |
| RP/Verifier       | [Source](https://github.com/trustbloc/sandbox) | [k8s](https://github.com/trustbloc/sandbox/tree/59b6dee1552f7afdea3a8b6f804e3dfdfcc0837e/k8s/rp)            |
| ORY Hydra         | [Source](https://github.com/ory/hydra)         | [k8s](https://github.com/trustbloc/sandbox/tree/59b6dee1552f7afdea3a8b6f804e3dfdfcc0837e/k8s/login-consent) |
| Login and Consent | [Source](https://github.com/trustbloc/sandbox) | [k8s](https://github.com/trustbloc/sandbox/tree/59b6dee1552f7afdea3a8b6f804e3dfdfcc0837e/k8s/login-consent) |
| ORY Oathkeeper    | [Source](https://github.com/ory/oathkeeper)    | [k8s](https://github.com/trustbloc/sandbox/tree/59b6dee1552f7afdea3a8b6f804e3dfdfcc0837e/k8s/cms)           |
| Strapi            | [Source](https://github.com/strapi/strapi)     | [k8s](https://github.com/trustbloc/sandbox/tree/59b6dee1552f7afdea3a8b6f804e3dfdfcc0837e/k8s/cms)           |


## v0.1.6
### Architecture Diagram
![VCS Architecture diagram v0.1.6](../images/vcs_component_diagram_v0.1.6.svg)

### Component Details
#### Core
| Component                   | Source URL                                                                 | Sample Docker Configuration                                                                                          |
|-----------------------------|----------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------|
| Router                      | [Source](https://github.com/trustbloc/hub-router)                          | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.6/test/bdd/fixtures/demo/docker-compose-wallet.yml)          |
| EDV                         | [Source](https://github.com/trustbloc/edv)                                 | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.6/test/bdd/fixtures/demo/docker-compose-edv.yml)             |
| EDV Oathkeeper              | [Source](https://github.com/ory/oathkeeper)                                | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.6/test/bdd/fixtures/demo/docker-compose-edv.yml)             |
| Ops KMS                     | [Source](https://github.com/trustbloc/hub-kms)                             | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.6/test/bdd/fixtures/demo/docker-compose-kms.yml)             |
| Ops KMS Oathkeeper          | [Source](https://github.com/ory/oathkeeper )                               | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.6/test/bdd/fixtures/demo/docker-compose-kms.yml)             |
| Authz KMS                   | [Source](https://github.com/trustbloc/hub-kms)                             | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.6/test/bdd/fixtures/demo/docker-compose-kms.yml)             |
| Authz KMS Oathkeeper        | [Source](https://github.com/ory/oathkeeper )                               | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.6/test/bdd/fixtures/demo/docker-compose-kms.yml)             |
| TrustBloc SignIn (Hub-Auth) | [Source](https://github.com/trustbloc/hub-auth)                            | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.6/test/bdd/fixtures/demo/docker-compose-auth.yml)            |
| TrustBloc SignIn (Hydra)    | [Source](https://github.com/ory/hydra)                                     | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.6/test/bdd/fixtures/demo/docker-compose-auth.yml)            |
| DID Resolver                | [Source](https://github.com/trustbloc/edge-service/tree/main/cmd/did-rest) | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.6/test/bdd/fixtures/demo/docker-compose-resolver.yml)        |
| TrustBloc DID Method        | [Source](https://github.com/trustbloc/trustbloc-did-method)                | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.6/test/bdd/fixtures/demo/docker-compose-did-method.yml)      |
| Sidetree Fabric             | [Source](https://github.com/trustbloc/sidetree-fabric)                     | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.6/test/bdd/fixtures/demo/docker-compose-sidetree-fabric.yml) |

#### Wallet
| Component     | Source URL                                        | Sample Docker Configuration                                                                                 |
|---------------|---------------------------------------------------|-------------------------------------------------------------------------------------------------------------|
| Wallet Web    | [Source](https://github.com/trustbloc/edge-agent) | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.6/test/bdd/fixtures/demo/docker-compose-wallet.yml) |
| Wallet Server | [Source](https://github.com/trustbloc/edge-agent) | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.6/test/bdd/fixtures/demo/docker-compose-wallet.yml) |

#### Services
| Component           | Source URL                                          | Sample Docker Configuration                                                                              |
|---------------------|-----------------------------------------------------|----------------------------------------------------------------------------------------------------------|
| Issuer VC Service   | [Source](https://github.com/trustbloc/edge-service) | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.6/test/bdd/fixtures/demo/docker-compose-vcs.yml) |
| Verifier VC Service | [Source](https://github.com/trustbloc/edge-service) | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.6/test/bdd/fixtures/demo/docker-compose-vcs.yml) |

#### Demo
| Component         | Source URL                                     | Sample Docker Configuration                                                                                            |
|-------------------|------------------------------------------------|------------------------------------------------------------------------------------------------------------------------|
| Issuer            | [Source](https://github.com/trustbloc/sandbox) | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.6/test/bdd/fixtures/demo/docker-compose-demo-applications.yml) |
| RP/Verifier       | [Source](https://github.com/trustbloc/sandbox) | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.6/test/bdd/fixtures/demo/docker-compose-demo-applications.yml) |
| ORY Hydra         | [Source](https://github.com/ory/hydra)         | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.6/test/bdd/fixtures/demo/docker-compose-auth.yml)              |
| Login and Consent | [Source](https://github.com/trustbloc/sandbox) | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.6/test/bdd/fixtures/demo/docker-compose-auth.yml)              |
| ORY Oathkeeper    | [Source](https://github.com/ory/oathkeeper)    | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.6/test/bdd/fixtures/demo/docker-compose-cms.yml)               |
| Strapi            | [Source](https://github.com/strapi/strapi)     | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.6/test/bdd/fixtures/demo/docker-compose-cms.yml)               |

## v0.1.5
### Architecture Diagram
![VCS Architecture diagram v0.1.5](../images/vcs_component_diagram_v0.1.5.svg)

### Component Details
#### Core
| Component                   | Source URL                                                                   | Sample Docker Configuration                                                                                                         |
|-----------------------------|------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------|
| Router                      | [Source](https://github.com/trustbloc/hub-router)                            | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.5/test/bdd/fixtures/demo/docker-compose-didcomm.yml#L11-L41)           |
| EDV                         | [Source](https://github.com/trustbloc/edv)                                   | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.5/test/bdd/fixtures/demo/docker-compose-edge-components.yml#L168-L187) |
| EDV Oathkeeper              | [Source](https://github.com/ory/oathkeeper)                                  | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.5/test/bdd/fixtures/demo/docker-compose-third-party.yml#L91-L112)      |
| Ops KMS                     | [Source](https://github.com/trustbloc/hub-kms)                               | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.5/test/bdd/fixtures/demo/docker-compose-edge-components.yml#L306-L339) |
| Ops KMS Oathkeeper          | [Source](https://github.com/ory/oathkeeper )                                 | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.5/test/bdd/fixtures/demo/docker-compose-third-party.yml#L114-L135)     |
| Authz KMS                   | [Source](https://github.com/trustbloc/hub-kms)                               | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.5/test/bdd/fixtures/demo/docker-compose-edge-components.yml#L267-L304) |
| Authz KMS Oathkeeper        | [Source](https://github.com/ory/oathkeeper )                                 | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.5/test/bdd/fixtures/demo/docker-compose-third-party.yml#L91-L112)      |
| TrustBloc SignIn (Hub-Auth) | [Source](https://github.com/trustbloc/hub-auth)                              | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.5/test/bdd/fixtures/demo/docker-compose-edge-components.yml#L168-L187) |
| TrustBloc SignIn (Hydra)    | [Source](https://github.com/ory/hydra)                                       | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.5/test/bdd/fixtures/demo/docker-compose-edge-components.yml#L381-L410) |
| DID Resolver                | [Source](https://github.com/trustbloc/edge-service/tree/main/cmd/did-rest)   | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.5/test/bdd/fixtures/demo/docker-compose-edge-components.yml#L150-L166) |
| TrustBloc DID Method        | [Source](https://github.com/trustbloc/trustbloc-did-method)                  | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.5/test/bdd/fixtures/demo/docker-compose-edge-components.yml#L251-L265) |
| Sidetree Fabric             | [Source](https://github.com/trustbloc/sidetree-fabric)                       | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.5/test/bdd/fixtures/demo/docker-compose-sidetree-fabric.yml)           |

#### Wallet
| Component     | Source URL                                        | Sample Docker Configuration                                                                                                         |
|---------------|---------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------|
| Wallet Web    | [Source](https://github.com/trustbloc/edge-agent) | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.5/test/bdd/fixtures/demo/docker-compose-edge-components.yml#L225-L249) |
| Wallet Server | [Source](https://github.com/trustbloc/edge-agent) | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.5/test/bdd/fixtures/demo/docker-compose-edge-components.yml#L189-L223) |

#### Services
| Component           | Source URL                                          | Sample Docker Configuration                                                                                                       |
|---------------------|-----------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------|
| Issuer VC Service   | [Source](https://github.com/trustbloc/edge-service) | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.5/test/bdd/fixtures/demo/docker-compose-edge-components.yml#L11-L43) |
| Verifier VC Service | [Source](https://github.com/trustbloc/edge-service) | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.5/test/bdd/fixtures/demo/docker-compose-edge-components.yml#L45-L77) |

#### Demo
| Component         | Source URL                                          | Sample Docker Configuration                                                                                                         |
|-------------------|-----------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------|
| Issuer            | [Source](https://github.com/trustbloc/sandbox)      | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.5/test/bdd/fixtures/demo/docker-compose-demo-applications.yml#L11-L45) |
| RP/Verifier       | [Source](https://github.com/trustbloc/sandbox)      | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.5/test/bdd/fixtures/demo/docker-compose-demo-applications.yml#L47-L71) |
| ORY Hydra         | [Source](https://github.com/ory/hydra)              | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.5/test/bdd/fixtures/demo/docker-compose-third-party.yml#L11-L40)       |
| Login and Consent | [Source](https://github.com/trustbloc/sandbox)      | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.5/test/bdd/fixtures/demo/docker-compose-third-party.yml#L239-L253)     |
| ORY Oathkeeper    | [Source](https://github.com/ory/oathkeeper)         | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.5/test/bdd/fixtures/demo/docker-compose-third-party.yml#L42-L63)      |
| Strapi            | [Source](https://github.com/strapi/strapi)          | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.5/test/bdd/fixtures/demo/docker-compose-third-party.yml#L137-L158)     |

## v0.1.4
### Architecture Diagram
![VCS Architecture diagram v0.1.4](../images/vcs_component_diagram_v0.1.4.svg)

### Component Details
#### Core
| Component            | Source URL                                                                   | Sample Docker Configuration                                                                                                         |
|----------------------|------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------|
| EDV                  | [Source](https://github.com/trustbloc/edv)                                   | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.4/test/bdd/fixtures/demo/docker-compose-edge-components.yml#L164-L177) |
| DID Resolver         | [Source](https://github.com/trustbloc/edge-service/tree/main/cmd/did-rest)   | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.4/test/bdd/fixtures/demo/docker-compose-edge-components.yml#L146-L162) |
| TrustBloc DID Method | [Source](https://github.com/trustbloc/trustbloc-did-method)                  | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.4/test/bdd/fixtures/demo/docker-compose-edge-components.yml#L198-L212) |
| Sidetree Fabric      | [Source](https://github.com/trustbloc/sidetree-fabric)                       | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.4/test/bdd/fixtures/demo/docker-compose-sidetree-fabric.yml)           |

#### Wallet
| Component  | Source URL                                        | Sample Docker Configuration                                                                                                         |
|------------|---------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------|
| User Agent | [Source](https://github.com/trustbloc/edge-agent) | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.4/test/bdd/fixtures/demo/docker-compose-edge-components.yml#L179-L196) |

#### Services
| Component           | Source URL                                          | Sample Docker Configuration                                                                                                       |
|---------------------|-----------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------|
| Issuer VC Service   | [Source](https://github.com/trustbloc/edge-service) | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.4/test/bdd/fixtures/demo/docker-compose-edge-components.yml#L11-L42) |
| Verifier VC Service | [Source](https://github.com/trustbloc/edge-service) | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.4/test/bdd/fixtures/demo/docker-compose-edge-components.yml#L44-L75) |

#### Demo
| Component         | Source URL                                          | Sample Docker Configuration                                                                                                         |
|-------------------|-----------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------|
| Issuer            | [Source](https://github.com/trustbloc/sandbox)      | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.4/test/bdd/fixtures/demo/docker-compose-demo-applications.yml#L11-L43) |
| RP/Verifier       | [Source](https://github.com/trustbloc/sandbox)      | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.4/test/bdd/fixtures/demo/docker-compose-demo-applications.yml#L45-L69) |
| ORY Hydra         | [Source](https://github.com/ory/hydra)              | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.4/test/bdd/fixtures/demo/docker-compose-third-party.yml#L11-L40)       |
| Login and Consent | [Source](https://github.com/trustbloc/sandbox)      | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.4/test/bdd/fixtures/demo/docker-compose-third-party.yml#L157-L171)     |
| ORY Oathkeeper    | [Source](https://github.com/ory/oathkeeper)         | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.4/test/bdd/fixtures/demo/docker-compose-third-party.yml#L42-L63)       |
| Strapi            | [Source](https://github.com/strapi/strapi)          | [Docker](https://github.com/trustbloc/sandbox/blob/v0.1.4/test/bdd/fixtures/demo/docker-compose-third-party.yml#L65-L86)       |
