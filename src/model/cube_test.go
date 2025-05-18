package model

import (
	"testing"
)

const StartCubeString = `[
  [
    [
      {
        "back": "yellow",
        "down": "green",
        "left": "red"
      },
      {
        "back": "yellow",
        "down": "green"
      },
      {
        "back": "yellow",
        "down": "green",
        "right": "orange"
      }
    ],
    [
      {
        "back": "yellow",
        "left": "red"
      },
      {
        "back": "yellow"
      },
      {
        "back": "yellow",
        "right": "orange"
      }
    ],
    [
      {
        "back": "yellow",
        "left": "red",
        "up": "blue"
      },
      {
        "back": "yellow",
        "up": "blue"
      },
      {
        "back": "yellow",
        "right": "orange",
        "up": "blue"
      }
    ]
  ],
  [
    [
      {
        "down": "green",
        "left": "red"
      },
      {
        "down": "green"
      },
      {
        "down": "green",
        "right": "orange"
      }
    ],
    [
      {
        "left": "red"
      },
      {},
      {
        "right": "orange"
      }
    ],
    [
      {
        "left": "red",
        "up": "blue"
      },
      {
        "up": "blue"
      },
      {
        "right": "orange",
        "up": "blue"
      }
    ]
  ],
  [
    [
      {
        "down": "green",
        "front": "white",
        "left": "red"
      },
      {
        "down": "green",
        "front": "white"
      },
      {
        "down": "green",
        "front": "white",
        "right": "orange"
      }
    ],
    [
      {
        "front": "white",
        "left": "red"
      },
      {
        "front": "white"
      },
      {
        "front": "white",
        "right": "orange"
      }
    ],
    [
      {
        "front": "white",
        "left": "red",
        "up": "blue"
      },
      {
        "front": "white",
        "up": "blue"
      },
      {
        "front": "white",
        "right": "orange",
        "up": "blue"
      }
    ]
  ]
]`

func TestToReadableJSON_Example(t *testing.T) {
	cube := NewCube()
	jsonStr, err := cube.ToReadableJSON()
	if err != nil {
		t.Fatalf("ToReadableJSON failed: %v", err)
	}
	if jsonStr != StartCubeString {
		t.Errorf("Expected %s, got %s", StartCubeString, jsonStr)
	}
}
