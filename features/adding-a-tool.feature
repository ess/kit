Feature: Adding a tool
  I want to be able to actually set up some tools, so I should have the
  ability to add one.

  Scenario: Successfully adding a tool
    Given there's not a blah tool configured
    When I run `kit add jq -image docker.io/wayneeseguin/c3tk -stream`
