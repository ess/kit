Feature: Showing tool info
  I can add tools and list tools, but it would also be neat if I could see
  more detailed information about the tools that I've added.

  Background:
    Given there's a jq tool available
    But there's not a blah tool available

  Scenario: Successfully getting tool information
    When I run `kit info jq`
    Then I see the configuration for the jq tool

    @failure
  Scenario: Tool name is a required argument
    When I run `kit info`
    Then I'm advised that I must provide a tool name
    And I see the info usage text

    @failure
  Scenario: Cannot see info for an unknown tool
    When I run `kit info blah`
    Then I'm advised that there is no blah tool
