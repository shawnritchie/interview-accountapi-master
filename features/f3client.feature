Feature: Form 3 Organisation - Account Client
  The following feature file describes the different scenarios
  that are to be handled within the Form 3 client. These include
  the Creation, Fetching, List, Patch & Delete functionality. The
  following test will both test the client side and server side
  validation ensuring safety at both ends.

  Scenario Outline: creating minimal accounts from various countries
    Given an initiated Client
    And a random organisationId
    And a random accountId
    When I send "POST" request to "/v1/organisation/accounts"
      | key               | value            |
      | Country           | <Country>        |
      | BankId            | <BankId>         |
      | BIC               | <BIC>            |
      | BankIdCode        | <BankIdCode>     |
      | AccountNumber     | <AccountNumber>  |
      | IBAN              | <IBAN>           |
      | Classification    | <Classification> |
    Then we expect a valid response
#    And response should contain an account Number
#    And response should contain an IBAN Number
  Examples:
    | Country | BankId      | BIC      | BankIdCode | AccountNumber | IBAN | Classification |
    | GB      | 000006      | NWBKGB22 | GBDSC      |               |      | Personal       |
    | AU      |             | NWBKGB22 | AUBSB      |               |      | Business       |
    | BE      | 000006      |          | BE         |               |      | Personal       |
    | CA      |             | NWBKGB22 | CACPA      |               |      | Business       |
    | FR      | 0000000010  |          | FR         |               |      | Personal       |
    | DE      | 00000008    |          | DEBLZ      |               |      | Business       |
    | GR      | 0000007     |          | GRBIC      |               |      | Personal       |
    | HK      |             | NWBKGB22 | HKNCC      |               |      | Business       |
    | IT      | 0000000010  |          | ITNCC      |               |      | Personal       |
    | IT      | 00000000011 |          | ITNCC      | 000000000012  |      | Business       |
    | LU      | 003         |          | LULUX      |               |      | Personal       |
    | NL      |             | NWBKGB22 |            |               |      | Business       |
    | PL      | 00000008    |          | PLKNR      |               |      | Personal       |
    | PT      | 00000008    |          | PTNCC      |               |      | Business       |
    | ES      | 00000008    |          | ESNCC      |               |      | Personal       |
    | CH      | 00005       |          | CHBCC      |               |      | Business       |
    | US      | 000000009   | NWBKGB22 | USABA      |               |      | Personal       |

  Scenario Outline: creating corrupted accounts and expecting failure
    Given an initiated Client
    And a random organisationId
    And a random accountId
    When I send "POST" request to "/v1/organisation/accounts"
      | key               | value            |
      | Country           | <Country>        |
      | BankId            | <BankId>         |
      | BIC               | <BIC>            |
      | BankIdCode        | <BankIdCode>     |
      | AccountNumber     | <AccountNumber>  |
      | IBAN              | <IBAN>           |
      | Classification    | <Classification> |
    Then we expect a bad request response
    Examples:
      | Country | BankId      | BIC      | BankIdCode | AccountNumber | IBAN                    | Classification |
      | GB      | 000006      | NWBKGB22 | GBDSC      |               | 00000000000000000020    | Personal       |
#      | GB      | 00005       | NWBKGB22 | GBDSC      |               |                         | Personal       |
#      | GB      | 0000007     | NWBKGB22 | GBDSC      |               |                         | Personal       |
#      | GB      | 000006      |          | GBDSC      |               |                         | Personal       |
#      | GB      | 000006      | NWBKGB22 | AUBSB      |               |                         | Personal       |
#      | GB      | 000006      | NWBKGB22 | GBDSC      | 0004          |                         | Personal       |
#      | GB      | 000006      | NWBKGB22 | GBDSC      | 0000000010    |                         | Personal       |
#      | AU      |             | NWBKGB22 | AUBSB      | 0004          |                         | Business       |
#      | AU      |             | NWBKGB22 | AUBSB      | 000000000012  |                         | Business       |
#      | CA      |             | NWBKGB22 | CACPA      |               | GB33BUKB20201555555555  | Business       |
#      | IT      | 00000000011 |          | ITNCC      |               |                         | Personal       |
#      | IT      | 0000000010  |          | ITNCC      | 000000000012  |                         | Business       |
#      | IT      | 00000000011 |          | NOTIT      | 000000000012  |                         | Business       |
#      | NL      | 0000000010  | NWBKGB22 |            |               |                         | Business       |
#      | NL      |             | NWBKGB22 | NOTNL      |               |                         | Business       |

  Scenario Outline: creating an account with names, alternative names and secondary identification
    Given an initiated Client
    And a random organisationId
    And a random accountId
    When I send "POST" request to "/v1/organisation/accounts"
      | key                       | value            |
      | Country                   | <Country>        |
      | BankId                    | <BankId>         |
      | BIC                       | <BIC>            |
      | BankIdCode                | <BankIdCode>     |
      | AccountNumber             | <AccountNumber>  |
      | IBAN                      | <IBAN>           |
      | Classification            | <Classification> |
      | Name                      | Shawn            |
      | Name                      | Stefanel         |
      | Name                      | Ritchie          |
      | Name                      | Extra            |
      | AlternativeNames          | Alternative      |
      | AlternativeNames          | Name             |
      | SecondaryIdentification   | Secondary        |
      | Status                    | confirmed        |
    Then we expect a valid response
    Examples:
      | Country | BankId      | BIC      | BankIdCode | AccountNumber | IBAN | Classification |
      | GB      | 000006      | NWBKGB22 | GBDSC      |               |      | Personal       |

#  Scenario Outline: creating an account with too many names
#    Given an initiated Client
#    And a random organisationId
#    And a random accountId
#    When I send "POST" request to "/v1/organisation/accounts"
#      | key                       | value            |
#      | Country                   | <Country>        |
#      | BankId                    | <BankId>         |
#      | BIC                       | <BIC>            |
#      | BankIdCode                | <BankIdCode>     |
#      | AccountNumber             | <AccountNumber>  |
#      | IBAN                      | <IBAN>           |
#      | Classification            | <Classification> |
#      | Name                      | Shawn            |
#      | Name                      | Stefanel         |
#      | Name                      | Ritchie          |
#      | Name                      | Extra            |
#      | Name                      | TooMany          |
#      | AlternativeNames          | Alternative      |
#      | AlternativeNames          | Name             |
#      | SecondaryIdentification   | Secondary        |
#      | Status                    | confirmed        |
#    Then we expect a bad request response
#    Examples:
#      | Country | BankId      | BIC      | BankIdCode | AccountNumber | IBAN | Classification |
#      | GB      | 000006      | NWBKGB22 | GBDSC      |               |      | Personal       |

#  Scenario Outline: creating an account with invalid status
#    Given an initiated Client
#    And a random organisationId
#    And a random accountId
#    When I send "POST" request to "/v1/organisation/accounts"
#      | key                       | value            |
#      | Country                   | <Country>        |
#      | BankId                    | <BankId>         |
#      | BIC                       | <BIC>            |
#      | BankIdCode                | <BankIdCode>     |
#      | AccountNumber             | <AccountNumber>  |
#      | IBAN                      | <IBAN>           |
#      | Classification            | <Classification> |
#      | Name                      | Shawn            |
#      | Name                      | Stefanel         |
#      | Name                      | Ritchie          |
#      | Name                      | Extra            |
#      | AlternativeNames          | Alternative      |
#      | AlternativeNames          | Name             |
#      | SecondaryIdentification   | Secondary        |
#      | Status                    | CORRUPTED        |
#    Then we expect a bad request response
#    Examples:
#      | Country | BankId      | BIC      | BankIdCode | AccountNumber | IBAN | Classification |
#      | GB      | 000006      | NWBKGB22 | GBDSC      |               |      | Personal       |

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