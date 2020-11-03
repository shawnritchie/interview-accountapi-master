# About me
I m Shawn and have been programming for way too long, to the degree that I feel that it is part of my identity. I love everything tech, I listen podcast weekly, read technical books, I use to attend conferences prior the pandemic, trying out new technologies!

Prior to this code challenge the last time I wrote go was around 3-4 years ago were I tried the language with some super simple examples to try out the concurrency model. In my regular day to day I am usually found coding in Java/Kotlin.

Overall I really enjoyed the challenge and got a better feel of the go language, and its unique traits.

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
- Do further testing especially on the Request method for each of the API as it wasn't tested as extensively as the UnsafeRequest

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

##Building

```
docker build -t shawnritchie/interview-accountapi-master:1.0 .
docker run shawnritchie/interview-accountapi-master:1.0
```