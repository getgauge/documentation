# Living and Executable Documentation

## Living Documentation

> A *living document* is a document that is continually edited and updated. - [wikipedia](https://en.wikipedia.org/wiki/Living_document)

In Gauge ecosystem, this provides information that is current, accurate and easy to understand.

Specs are written in a natural language - Markdown format. Each spec describes how a particular piece of functionality is supposed to behave, gives various scenarios and describes the desired outcome.

Business stakeholders can review the documentation to ensure that it describes the desired behavior of the system from a end-user standpoint. Developers can use the information to help them program only what is needed, making the code as lean as possible. Testers can use the documentation to implement/create test cases and confirm that results meet requirements, and Operations can rely on the same information to support ongoing maintenance operations in the production environment.

## Executable documentation

Getting effective feedback is crucial in creating the executable specification. Running stable automated tests that cover the functionality in CI gives the required feedback.

A change in requirement would trigger a change in the specification and the corresponding underlying test implementation documenting the new functionality. If the checks are verified periodically, the new change in requirement could result in a failure when the change is yet to be implemented in the system, or vice versa. Fixing a failing build would serve as a mark of the completion of the functionality. Such documents are, thus, a reliable source of information on business functionality of underlying software.

Unlike traditional static documentation which can get outdated, executable living documentation could be a complete set of specifications with checks for correctness of the implementations.
