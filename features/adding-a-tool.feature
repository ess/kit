Feature: Adding a tool
  I want to be able to actually set up some tools, so I should have the
  ability to add one.

  Scenario: Successfully adding a tool with defaults
    Given there's not a jq tool configured
    When I run `kit add jq`
    Then the jq tool is configured
    And a jq symlink to kit now exists
    And jq's image is docker.io/wayneeseguin/c3tk
    And jq's default tag is latest
    And jq is set up to stream IO
    But jq is not set up with a TTY

  Scenario Outline: Successfully adding a tool with a specific image
    Given there's not a jq tool configured
    When I run `kit add jq <Flag> ess/c3tk`
    Then the jq tool is configured
    And jq's image is docker.io/ess/c3tk
    And that is jq's only non-default setting

    Examples:
      | Flag    |
      | -i      |
      | --image |

  Scenario: Successfully adding a tool with a specific tag
    Given there's not a jq tool configured
    When I run `kit add jq --tag variant`
    Then the jq tool is configured
    And jq's tag is variant
    And that is jq's only non-default setting

  Scenario Outline: Successfully adding a tool with a tty
    Given there's not a jq tool configured
    When I run `kit add jq <Flag>`
    Then the jq tool is configured
    And jq's tty is true
    And that is jq's only non-default setting

    Examples:
      | Flag  |
      | -t    |
      | --tty |

  Scenario: Successfully adding a tool without streaming
    Given there's not a jq tool configured
    When I run `kit add jq <Flag> variant`
    Then the jq tool is configured
    And jq's stream is false
    And that is jq's only non-default setting

    Examples:
      | Flag        |
      | -n          |
      | --no-stream |

  Scenario Outline: Updating a tool
    Given there is already a jq tool configured with default settings
    When I run `kit add jq <Flag>`
    Then jq's configuration is updated

    Examples:
      | Flag                    |
      | -i whatever/image       |
      | --image whatever/image  |
      | --tag sometag           |
      | -t                      |
      | --tty                   |
      | -n                      |
      | --no-stream             |
