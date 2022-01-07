Feature: Adding a tool
  I want to be able to actually set up some tools, so I should have the
  ability to add one.

  Background:
    Given there's not a jq tool configured

  Scenario: Successfully adding a tool with default behavior
    When I run `kit add jq`
    Then the jq tool is configured
    And a jq symlink to kit now exists
    And jq's image is docker.io/wayneeseguin/c3tk
    And jq's image gets pulled from upstream
    And jq's default tag is latest
    And jq is set up to stream IO

  Scenario Outline: Successfully adding a tool with a specific image
    When I run `kit add jq <Flag> ess/c3tk`
    Then the jq tool is successfully added with the docker.io/ess/c3tk image

    Examples:
      | Flag    |
      | -i      |
      | --image |

  Scenario Outline: Successfully adding a tool with a specific tag
    When I run `kit add jq <Flag> variant`
    Then the jq tool is successfully added with the variant tag

    Examples:
      | Flag  |
      | -t    |
      | --tag |

  Scenario Outline: Successfully adding a tool without streaming
    When I run `kit add jq <Flag> variant`
    Then the jq tool is successfully added with IO streaming disabled

    Examples:
      | Flag        |
      | -n          |
      | --no-stream |

  Scenario Outline: Updating a tool
    Given there is already a jq tool configured with default settings
    When I run `kit add jq <Flag>`
    Then jq's configuration is updated
    And jq's image gets pulled from upstream

    Examples:
      | Flag                    |
      | -i whatever/image       |
      | --image whatever/image  |
      | -t sometag              |
      | --tag sometag           |
      | -n                      |
      | --no-stream             |
