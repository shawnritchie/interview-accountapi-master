# Form3 Take Home Exercise

## Instructions
The goal of this exercise is to write a client library 
in Go to access our fake [account API](http://api-docs.form3.tech/api.html#organisation-accounts) service. 

### Should
- Client library should be written in Go
- Document your technical decisions
- Implement the `Create`, `Fetch`, `List` and `Delete` operations on the `accounts` resource. Note that filtering of the List operation is not required, but you should support paging
- Ensure your solution is well tested to the level you would expect in a commercial environment. Make sure your tests are easy to read.
- To keep this exercise simple, fields `data.attributes.private_identification`, `data.attributes.organisation_identification` 
and `data.relationships` were omitted in the provided fake accountapi implementation - do not implement these in your model
- If you encounter any problems running the fake accountapi we would encourage you to do some debugging first, 
before reaching out for help

#### Docker-compose
 - Add your solution to the provided docker-compose file
 - We should be able to run `docker-compose up` and see your tests run against the provided account API service 

### Please don't
- Use a code generator to write the client library
- Use a library for your client (e.g: go-resty). Only test libraries are allowed.
- Implement an authentication scheme

## How to submit your exercise
- Include your name in the README. If you are new to Go, please also mention this in the README so that we can consider this when reviewing your exercise
- Create a private [GitHub](https://help.github.com/en/articles/create-a-repo) repository, copy the `docker-compose` from this repository
- [Invite](https://help.github.com/en/articles/inviting-collaborators-to-a-personal-repository) @form3tech-interviewer-1 to your private repo
- Let us know you've completed the exercise using the link provided at the bottom of the email from our recruitment team

## License
Copyright 2019-2020 Form3 Financial Cloud

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.


## About me

## Architectural Decisions

---

##### 1. Linter - Go Lang CI Lint & ErrCheck

Seeing I'm somewhat of a noobie, to go, I need hand holding, so I introduced the usage of a linter & errorCheck, 
for hand holding making sure I follow community conventions and not missing out on unchecked errors

###### Setup
Go Lang CI Lint
```
brew install golangci/tap/golangci-lint
brew upgrade golangci/tap/golangci-lint
golangci-lint run
```

ErrCheck
```
go get -u github.com/kisielk/errcheck
errcheck .
```

---

##### 2. Why write a client library vs writing custom http client per dependency**

- Pro
    - Client library could avoid the repetition of code, especially if 
    the api is extensive and complex this could increase development lead 
    time and hand hold engineers in its proper usage
- Cons
    - Version could create a headache if left unattended too whereby multiple
    clients could be using different versions of the client library. This could 
    be resolved with proper version and making sure backward compatibility is 
    prioritised per release.
    - Unnecessary coupling between the server and client as the consuming service will
      depend on the entire payloads rather than the required fields
        

##### 3. How will we validate the functionality of the client
- _Client side validation of request against requirements/documentation_
    - Cons
        - Validation http request payload will couple the requirements of the client with the server,
          this will imply that changes in the functionality will be undetected with validation
    - Pro
        - An opinionated client library will reduce the burden from the engineer to look into the documentation
          and will allow them to use an opinionated client library and find discrepancies quickly through unit test
    - Alternative
        - Implement opinionated builders but keep the flexibility inside the client to allow the engineer to
          think outside the box
- _Alternative implement tests against server requirements_
    - Cons
        - Without proper schema/contract tooling we will need to relay on end to end testing which
          could be considered slow, and tedious to write as we are relaying on a full environment, 
          this could potentially break engineers flow. On the upside these are very reliable and 
          holistic tests which could identify potential bugs not just regression in the schema.
    - Pro
        - each released version will test for regression in the server's api which could lead to 
        potential bugs
    - Alternative
        - Introduce a schema registry or contract testing with a centralised contract server this 
        way we can avoid the end to end testing or blackbox test with lightweight unit test which 
        validate schema/contracts between producer and consumer
    
    
## Issues Found
The following issues have been found while testing the client against the black box image of the account api.
Test proving the issues have been provided but have been commited out to make sure the build runs sucessfully.

### Create
1. No Validation on BankID
2. No validation on the number of names
3. No validation on the status field
4. No validation mapping the country requirements specified in the documentation highlight in the cucumber test
5. The response is a mirror of the request parameters which does not match the specifications

### Fetch
6. No UUID validation on the ID field inside of the payload

### List
7. Ignores the page size parameter and constantly return 1000 records

### Delete
8. Returns a 204 when deleting a non existent an account
9. Deleting an account with the incorrect version number returns a 404 instead of a 409

## Future Improvements
- Validation of host environmental variable to make sure it matches the library expectations
- Research more indepth the use of channels in the context of how the client API has been designed seeing the current 
design makes use of one time use channel seeing it closes the channel once the results are posted.
- Test for race conditions / potential deadlock in the builders especially the list builder using shared state for page traversal

## Production Ready TODO
- [ ] Exponential Backoff https://api-docs.form3.tech/api.html#introduction-and-api-conventions-timeouts-rate-limiting-and-retry-strategy
- [ ] Circuit Breaking
- [ ] Health check integration and exposure
- [ ] Metrics 
- [ ] Distributed tracing integration
- [ ] Investigate automated alerting integration
- [ ] Configurable logging
- [ ] Potential contract testing 

## Clients Specs 
The following snippets show the client specification we would like to develop

### Create
```
F3Client
    .Create()
    .WithX()
    .WithY()
    .UnsafeRequest(context, response, errors)
    .Request(context, response, errors)
```

### Fetch
```
F3Client
    .Fetch()
    .WithX()
    .UnsafeRequest(context, response, errors)
    .Request(context, response, errors)
```    

### Delete
```
F3Client
    .Delete()
    .WithX()
    .UnsafeRequest(context, errors)
    .Request(context, errors)
```    

### List
```
F3Client
    .List()
    .WithPage(1)
    .WithSize(10)
    .UnsafeRequest(context, response, errors)
    .Request(context, response, errors)
    .Next()
    .Prev()
    .First
    .Last()
```    

    