# RingCentral Engage API Docs

[![Build Status][build-status-svg]][build-status-link]
[![Docs Status][docs-status-svg]][docs-status-link]
[![Docs][docs-svg]][docs-link]

This repository is the home of the RingCentral Engage Developer Guide: a collection of materials, and documentation to help educate developers on how to build on top of the RingCentral Engage platform.

Visit at: https://engage-api-docs.rtfd.org

## Contributing

If you would like to contribute to the RingCentral documentation effort, fork this repository, make your desired edits and contributions, and issue a pull request accordingly.

### Running ReadTheDocs Locally

```
% git clone https://github.com/ringcentral/engage-digital-api-docs.git
% cd engage-digital-api-docs
% pip install mkdocs
% pip install mkdocs-ringcentral
% mkdocs serve
```

Then you should be able to load http://localhost:8000 to view the documentation.

### Testing OpenAPI Specs

This repo holds OpenAPI specs for Engage Digital and Engage Voice. For each commit, tests are run on Travis CI to verify that the OpenAPI 3.0 specs validate.

* Engage Digital Spec: [specs/digital/engage-digital_openapi3.yaml](specs/digital/engage-digital_openapi3.yaml)
* Engage Voice Spec: [specs/voice/engage-voice_openapi3.json](specs/voice/engage-voice_openapi3.json)
* Tests: [specs_test.go](specs_test.go)

You can verify the specs localy with the following if you have [Go installed](https://golang.org/).

```
% git clone https://github.com/ringcentral/ringcentral-api-docs.git
% cd ringcentral-api-docs
% go test -v ./...
```

If you wish to change the specs being tested edit the [specs_test.go](specs_test.go) file.

 [docs-status-svg]: https://travis-ci.com/ringcentral/engage-api-docs.svg?branch=master
 [docs-status-link]: https://readthedocs.org/projects/engage-api-docs/builds/
 [build-status-svg]: https://travis-ci.com/ringcentral/engage-api-docs.svg?branch=master
 [build-status-link]: https://travis-ci.com/ringcentral/engage-api-docs
 [docs-svg]: https://img.shields.io/badge/docs-readthedocs-blue.svg
 [docs-link]: https://engage-api-docs.readthedocs.io/