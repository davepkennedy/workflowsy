Feature: Basic Calculator

Background:
    Given I have a calculator

Scenario Outline:
    When I input "<input>"
    Then the result should be <result>

    Examples:
    | input          | result |
    | 2 3 +          | 5      |
    | 2 3 -          | -1     |
    | 2 3 *          | 6      |
    | 6 3 /          | 2      |
    | 2 3 * 10 2 / * | 30     |