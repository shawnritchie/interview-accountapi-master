Feature: Form 3 Organisation - Account Client
  The following feature file describes the different scenarios
  that are to be handled within the Form 3 client specificly the.
  List api. The following test will both test the client side and
  server side validation ensuring safety at both ends.

  Scenario: get first page from list api
    Given an initiated Client
    And a random organisationId
    And a random accountId
    And a valid account has been registered
    When I send "GET" request to "/v1/organisation/accounts/" for page 1 with a page size of 1
    Then we expect a valid response
#    And a page size of 1

#  Scenario: get first page from list api
#    Given an initiated Client
#    And a random organisationId
#    And a random accountId
#    And a valid account has been registered
#    And I send "GET" request to "/v1/organisation/accounts/" for page 1 with a page size of 1
#    When I navigate to the next page
#    Then we expect a valid response
#    And a page size of 1

  Scenario: get previous page from list api
    Given an initiated Client
    And a random organisationId
    And a random accountId
    And a valid account has been registered
    And I send "GET" request to "/v1/organisation/accounts/" for page 2 with a page size of 1
    When I navigate to the previous page
    Then we expect a valid response
#    And a page size of 1

  Scenario: get first page from list api
    Given an initiated Client
    And a random organisationId
    And a random accountId
    And a valid account has been registered
    And I send "GET" request to "/v1/organisation/accounts/" for page 2 with a page size of 1
    When I navigate to the first page
    Then we expect a valid response
#    And a page size of 1

  Scenario: get last page from list api
    Given an initiated Client
    And a random organisationId
    And a random accountId
    And a valid account has been registered
    And I send "GET" request to "/v1/organisation/accounts/" for page 1 with a page size of 1
    When I navigate to the last page
    Then we expect a valid response
#    And a page size of 1

  Scenario: get first page from list api
    Given an initiated Client
    And a random organisationId
    And a random accountId
    And a valid account has been registered
    When I send "GET" request to "/v1/organisation/accounts/" for page 10000 with a page size of 1000
    Then we expect a valid response
    And a page size of 0