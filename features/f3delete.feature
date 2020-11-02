Feature: Form 3 Organisation - Account Client
  The following feature file describes the different scenarios
  that are to be handled within the Form 3 client specifically the
  Delete api. The following test will both test the client side and
  server side validation ensuring safety at both ends.

  Scenario: deleting an account successfully
    Given an initiated Client
    And a random organisationId
    And a random accountId
    And a valid account has been registered
    When I send "DELETE" request to "/v1/organisation/accounts/" with the accountId and version 0
    Then we expect a valid response

#  Scenario: deleting an account with incorrect version should return 409
#    Given an initiated Client
#    And a random organisationId
#    And a random accountId
#    And a valid account has been registered
#    When I send "DELETE" request to "/v1/organisation/accounts/" with the accountId and version 5
#    Then we expect a http status code: Conflict
#
#  Scenario: deleting an account a none existent account id
#    Given an initiated Client
#    And a random accountId
#    When I send "DELETE" request to "/v1/organisation/accounts/" with the accountId and version 0
#    Then we expect a http status code: Not Found