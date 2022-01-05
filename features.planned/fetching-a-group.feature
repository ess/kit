Feature: Fetching a group
  Predefined groups of tools are a thing, and I like convenience, so I'd like
  to be able to fetch those groups.

  Background:
    Given there is a kubernetes group in the default repo
    And there is a kubernetes group in https://example.com/repo
    And I have no groups configured

  Scenario: Fetching a group from the default repo
    When I run `kit fetch kubernetes`
    Then the group is fetched from the default repo
    And it is stored in my group configs location
    And the tools defined by that group are added

  Scenario Outline: Fetching a group from a specific repo
    When I run `kit fetch <Flag> https://example.com/repo kubernetes`
    Then the group is fetched from https://example.com/repo
    And it is stored in my group configs location
    And the tools defined by that group are added

    Examples:
      | Flag    |
      | -r      |
      | --repo  |

    @failure
  Scenario: Failing to fetch a group that does not exist
    Given there is not a kubernetes group in the default repo
    When I run `kit fetch kubernetes`
    Then I see an error regarding the inability to fetch the group
    And the group is not stored in my group configs location
    And no tools are added
