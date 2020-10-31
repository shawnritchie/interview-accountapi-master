Feature: Form 3 Organisation - Account Client
  The following feature file describes the different scenarios
  that are to be handled within the Form 3 client. These include
  the Creation, Fetching, List, Patch & Delete functionality. The
  following test will both test the client side and server side
  validation ensuring safety at both ends.

  Scenario Outline: creating minimal accounts from various countries
    Given an initiated Client
    And an organisationId of "634e3a41-26b8-49f9-a23d-26fa92061f38"
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
    And an organisationId of "634e3a41-26b8-49f9-a23d-26fa92061f38"
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

