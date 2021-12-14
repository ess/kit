Feature: Listing Tools
  Just so I know what my capabilities are, I want to be able to list the tools
  that are managed by kit.

  Scenario: Listing tools when there are none
    Given there are no tools configured
    When I run `kit list`
    Then there is no output

  Scenario: Listing tools when there are some
    Given there are tools configured
    When I run `kit list`
    Then there is output
    And it mentions each of the configured tools
