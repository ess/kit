Feature: Removing a tool
  So, I know how to add tools when I need them, but I also want to remove them
  when I no longer need them (because I've found a cool replacement or just
  don't want extra stuff laying around).

  Scenario: Removing a known tool
    Given the jq tool is available
    When I run `kit rm jq`
    Then the jq symlink is removed
    But there is no output
    And no other changes are made

  Scenario: Removing an unknown tool
    Given there are no tools available
    When I run `kit rm jq`
    Then there is no output
    And no changes are made
