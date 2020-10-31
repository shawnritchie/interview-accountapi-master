Feature: Form 3 Organisation - Account Builder
  The following feature file describes the different scenarios
  that are to be handled within the Form 3 client. The scope of the
  builder is to validate requests prior them being submitted, this should
  aid engineers finding issues locally using faster and cheaper tests.

  Scenario Outline: building payload for accounts from various countries
    When we validate the "POST" account builder with properties
      | key               | value            |
      | Country           | <Country>        |
      | BankId            | <BankId>         |
      | BIC               | <BIC>            |
      | BankIdCode        | <BankIdCode>     |
      | AccountNumber     | <AccountNumber>  |
      | IBAN              | <IBAN>           |
      | Classification    | <Classification> |
    Then we expect no validation errors
    Examples:
      | Country | BankId      | BIC      | BankIdCode | AccountNumber | IBAN | Classification |
      | GB      | 000006      | NWBKGB22 | GBDSC      |               |      | Personal       |
      | AU      |             | NWBKGB22 | AUBSB      |               |      | Business       |
      | BE      | 003         |          | BE         |               |      | Personal       |
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

  Scenario Outline: building payload for accounts from various countries
    When we validate the "POST" account builder with properties
      | key               | value            |
      | Country           | <Country>        |
      | BankId            | <BankId>         |
      | BIC               | <BIC>            |
      | BankIdCode        | <BankIdCode>     |
      | AccountNumber     | <AccountNumber>  |
      | IBAN              | <IBAN>           |
      | Classification    | <Classification> |
    Then we expect validation errors
    Examples:
      | Country | BankId      | BIC      | BankIdCode | AccountNumber | IBAN                    | Classification |
      | GB      | 000006      | NWBKGB22 | GBDSC      |               | 00000000000000000020    | Personal       |
      | GB      | 00005       | NWBKGB22 | GBDSC      |               |                         | Personal       |
      | GB      | 0000007     | NWBKGB22 | GBDSC      |               |                         | Personal       |
      | GB      | 000006      |          | GBDSC      |               |                         | Personal       |
      | GB      | 000006      | NWBKGB22 | AUBSB      |               |                         | Personal       |
      | GB      | 000006      | NWBKGB22 | GBDSC      | 0004          |                         | Personal       |
      | GB      | 000006      | NWBKGB22 | GBDSC      | 0000000010    |                         | Personal       |
      | AU      |             | NWBKGB22 | AUBSB      | 0004          |                         | Business       |
      | AU      |             | NWBKGB22 | AUBSB      | 000000000012  |                         | Business       |
      | CA      |             | NWBKGB22 | CACPA      |               | GB33BUKB20201555555555  | Business       |
      | IT      | 00000000011 |          | ITNCC      |               |                         | Personal       |
      | IT      | 0000000010  |          | ITNCC      | 000000000012  |                         | Business       |
      | IT      | 00000000011 |          | NOTIT      | 000000000012  |                         | Business       |
      | NL      | 0000000010  | NWBKGB22 |            |               |                         | Business       |
      | NL      |             | NWBKGB22 | NOTNL      |               |                         | Business       |

