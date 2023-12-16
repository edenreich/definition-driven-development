# Definition-driven Development

- [Definition-driven Development](#definition-driven-development)
  - [One-Off / Initialize a Project](#one-off--initialize-a-project)
  - [Development Process](#development-process)
  - [Customization](#customization)
  - [Validation](#validation)
  - [Todos](#todos)

It's not a new concept, yet it has gained more popularity over the last few years due to an improvements in existing tooling.

Definition-driven development is the approach where a clear and well-defined understanding of project requirements and goals guides the entire development process. This approach emphasizes the importance of creating a detailed specification or definition of the project before initiating any coding or implementation.

Having a comprehensive definition upfront enables developers to better understand what needs to be achieved, leading to more efficient and focused development efforts. This approach helps minimize misunderstandings, reduces the likelihood of scope changes, and provides a roadmap for the development team to follow. It is often associated with methodologies like Waterfall, where the development process is sequential and relies on a thorough upfront planning phase.

Parallel Development is one of the key benefits of this approach. With a well-established interface, client and server development can proceed concurrently. Different teams or developers can work on the client and server components simultaneously, making the development process more efficient and reducing unclarity.

Once the interface is well-defined, it's also possible to adjust specific requirements and focus on generating only the things that are frequently changing from that interface.

Relying on OAS as the source of truth enables us to leverage the latest OAS version, specifically version 3.0.0 and beyond. As of the time of writing, when employing tools such as Swaggo, which generates OAS from Go code using annotations, there is a constraint limiting compatibility to OAS version 2.0.0. Furthermore, being independent of the annotations tied to a particular programming language enhances our flexibility, making us language-agnostic.

This method is frequently referred to as the "Design-First" approach. By following this strategy, seamless collaboration among various stakeholders becomes easily achievable.

## One-Off / Initialize a Project

All you need is an OpenAPI Specification (OAS). Once you have it well-defined, you can generate the first `client` and `server` projects in any desired programming language using `openapi-generator-cli`. If you're familiar with `Docker`, you can view the commands I placed in the Makefile; those container images already provide the `openapi-generator-cli` and the tooling needed for generating the code.

After the boilerplate code has been generated, it's necessary to review and determine what is really needed and what is frequently changing. The things that are likely never to change should not be generated, so add them to the `.openapi-generator-ignore` file.

Make the adjustments to those files you don't want to generate.

## Development Process

Let's say a new business requirements arrive at the table. They need a new endpoint.

1. Check out a new branch from the `main` upstream.
2. Design the endpoint by adjusting the source of truth, the OAS (`pets_api.yaml`).
3. Run `make generate`. Which will run the `openapi-generator-cli` for generating the `client` and `server`
4. Fill in the todos and write the business logic / implementation.
5. Create a Pull Request / Merge Request.
6. Let your co-worker review.
7. Once approved, merge it in.

Let's say those who work on the `client` are on a different team than those who work on the `server`. In that scenario, you can still generate the code for the `client` even on a completely separate repository, but the interface (the `OpenAPI Specification`) has to be with the same version. This way, teams can first design the specification and then work in parallel.

## Customization

Every project comes with unique requirements, making this aspect likely the most challenging part of the equation. However, once you successfully establish the setup, the subsequent workflow and minimal adjustments required for implementation can significantly boost the development velocity.

Utilizing pre-existing templates empowers us to select the generated API code seamlessly. There are instances, though, where bespoke templates become necessary, and fortunately, it's entirely feasible to override the existing ones.

By overriding the default templates, we gain the ability to delve into specific implementation details, thereby enhancing the speed of development. This flexibility ensures a tailored approach that aligns precisely with the project's distinctive needs.

Let's modify an existing template:

1. First let's fetch the template:
```bash
make openapi ARGS='author template -g go-server -o templates/go-server'
```
2. Modify it.
3. Modify the generate command by adding `-t templates/go-server`.

A variety of additional templates crafted by the open-source community are available for selection:

- ada
- ada-server
- android
- apache2
- apex
- asciidoc
- aspnetcore
- avro-schema
- bash
- crystal
- c
- clojure
- cwiki
- cpp-qt-client
- cpp-qt-qhttpengine-server
- cpp-pistache-server
- cpp-restbed-server
- cpp-restbed-server-deprecated
- cpp-restsdk
- cpp-tiny
- cpp-tizen
- cpp-ue4
- csharp
- csharp-functions
- dart
- dart-dio
- eiffel
- elixir
- elm
- erlang-client
- erlang-proper
- erlang-server
- fsharp-functions
- fsharp-giraffe-server
- go
- go-echo-server
- go-server
- go-gin-server
- graphql-schema
- graphql-nodejs-express-server
- groovy
- kotlin
- kotlin-server
- kotlin-spring
- kotlin-vertx
- ktorm-schema
- haskell-http-client
- haskell
- haskell-yesod
- java
- jaxrs-cxf-client
- java-helidon-client
- java-helidon-server
- java-inflector
- java-micronaut-client
- java-micronaut-server
- java-msf4j
- java-pkmst
- java-play-framework
- java-undertow-server
- java-vertx
- java-vertx-web
- java-camel
- jaxrs-cxf
- jaxrs-cxf-extended
- jaxrs-cxf-cdi
- jaxrs-jersey
- jaxrs-resteasy
- jaxrs-resteasy-eap
- jaxrs-spec
- javascript
- javascript-apollo-deprecated
- javascript-flowtyped
- javascript-closure-angular
- jetbrains-http-client
- jmeter
- julia-client
- julia-server
- k6
- lua
- markdown
- mysql-schema
- n4js
- nim
- nodejs-express-server
- objc
- ocaml
- openapi
- openapi-yaml
- plantuml
- perl
- php
- php-nextgen
- php-laravel
- php-lumen
- php-slim4
- php-symfony
- php-mezzio-ph
- php-dt
- postman-collection
- powershell
- protobuf-schema
- python
- python-pydantic-v1
- python-fastapi
- python-flask
- python-aiohttp
- python-blueplanet
- r
- ruby
- ruby-on-rails
- ruby-sinatra
- rust
- rust-server
- scalatra
- scala-akka
- scala-akka-http-server
- scala-finch
- scala-gatling
- scala-lagom-server
- scala-play-server
- scala-sttp
- scala-sttp4
- scalaz
- spring
- dynamic-html
- html
- html2
- swift5
- swift-combine
- typescript
- typescript-angular
- typescript-aurelia
- typescript-axios
- typescript-fetch
- typescript-inversify
- typescript-jquery
- typescript-nestjs
- typescript-node
- typescript-redux-query
- typescript-rxjs
- wsdl-schema
- xojo-client
- zapier

## Validation

To validate the schema of he OAS, just run:
```bash
make openapi ARGS='validate -i pets_api.yaml' 
```

## Todos

- [ ] cover virtual services / mocked services
- [ ] generate metrics based of customisations attributes in the OAS