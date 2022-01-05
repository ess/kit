Feature: Configuring Tools
  Each tool is defined with a config file. Since it's possible to distribute
  those config files, I would like to able to do so and also have a meaningful
  way to absorb all of the tools for which I have a config.

  Scenario: Configuring all tools
    Given I have no tools available
    But I have a collection of tool configs

    When I place the tool configs into my config dir
    And I run `kit configure`

    Then all of the tools for which I added a config are available

  Scenario: Reconfiguring tools
    Given I have a set of tools available
    When I run `kit configure`
    Then all of my available tools are reconfigured

  Scenario: No tools
    Given I have no tools available
    When I run `kit configure`
    Then no changes are made
    And I see no output
    And there are no available tools
