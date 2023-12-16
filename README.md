# Definition-driven Development

- [Definition-driven Development](#definition-driven-development)
  - [One-Off / Initialize a Project](#one-off--initialize-a-project)
  - [Development Process](#development-process)
  - [Customization](#customization)
  - [Todos](#todos)

It's not a new concept, yet it has gained more popularity over the last few years due to an improvements in existing tooling.

Definition-driven development is the approach where a clear and well-defined understanding of project requirements and goals guides the entire development process. This approach emphasizes the importance of creating a detailed specification or definition of the project before initiating any coding or implementation.

Having a comprehensive definition upfront enables developers to better understand what needs to be achieved, leading to more efficient and focused development efforts. This approach helps minimize misunderstandings, reduces the likelihood of scope changes, and provides a roadmap for the development team to follow. It is often associated with methodologies like Waterfall, where the development process is sequential and relies on a thorough upfront planning phase.

Parallel Development is one of the key benefits of this approach. With a well-established interface, client and server development can proceed concurrently. Different teams or developers can work on the client and server components simultaneously, making the development process more efficient and reducing unclarity.

Once the interface is well-defined, it's also possible to adjust specific requirements and focus on generating only the things that are frequently changing from that interface.

## One-Off / Initialize a Project

All you need is an OpenAPI Specification (OAS). Once you have it well-defined, you can generate the first `client` and `server` projects in any desired programming language using `openapi-generator-cli`. If you're familiar with `Docker`, you can view the commands I placed in the Makefile; those container images already provide the `openapi-generator-cli` and the tooling needed for generating the code.

After the boilerplate code has been generated, it's necessary to review and determine what is really needed and what is frequently changing. The things that are likely never to change should not be generated, so add them to the `.openapi-generator-ignore` file.

Make the adjustments to those files you don't want to generate.

## Development Process

Let's say a new business requirements arrive at the table. They need a new endpoint.

1. Check out a new branch from the `main` upstream.
2. Design the endpoint by adjusting the souce of truth, the OAS (`cats_api.yaml`).
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

## Todos

- [ ] add example template and explanation how to modiy the generated code
