# Fleet Tracker

## Description

This is a full stack appplication cunstructed using Vue.js for the frontend architecture and Go for the backend. Minimal external UI packages are used aside from Vue.js, and all of the components are custom-made. A user can use any OneStepGPS API key to view and track their devices. All user preferences are stored on the backend using a JSON file, allowing for portability. The state of the site is saved in a single obejct allowing for site-wide updates (one component can update multiple others without much effort). State is managed through Vue.js reactive data, and seperate controllers handle updates.

## Installation

To install this application, open `terminal` and clone the repository by running:

    git clone https://github.com/benw10-1/OneStepGPS-Application.git

## Usage

Open the `terminal` and run:

    npm install
    npm run build
    cd server
    go run server.go

## Links

Visit the live site @ http://onestepgpsapp.ddns.net/
Or view the [demonstration video](https://youtu.be/yEiQoiYzoHw)
