#!/bin/bash
set -e

avalanche subnet delete lilypad
avalanche network stop
avalanche network clean
