Feature: Form 3 Organisation - Account Client
  The following feature file describes the different scenarios
  that are to be handled within the Form 3 client specifically the.
  Fetch api. The following test will both test the client side and
  server side validation ensuring safety at both ends.

  Scenario: fetching an account successfully
    Given an initiated Client
    And a random organisationId
    And a random accountId
    And a valid account has been registered
    When I send "GET" request to "/v1/organisation/accounts/" with the accountId
    Then we expect a valid response

#  Scenario: fetching an account with an invalid uuid
#    Given an initiated Client
#    And a random organisationId
#    And an invalid accountId
#    When I send "GET" request to "/v1/organisation/accounts/" with the accountId
#    Then we expect a validation error

  Scenario: fetching an account for an account which does not exist
    Given an initiated Client
    And a random organisationId
    And a random accountId
    When I send "GET" request to "/v1/organisation/accounts/" with the accountId
    Then we expect a http status code: Not Found