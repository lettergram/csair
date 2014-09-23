CS242 - Programming Studio - Assignment 2
========

Written by Austin Walters
Created Date: Spring 2014
Language: Golang

### Problem Description

You are a senior software engineer for a new international airline, CSAir. Before CSAir launches passenger service across the globe, they first need to start selling tickets to passengers. But before that can happen, some software needs to be written to manage the extensive route map. You have been tasked by CEO Woodley to begin work on this software. Specifically, the initial requirements for this new software are:

* Parse the raw data that represents CSAir's route map into a data structure in memory
* Allow users of the software to query data about each of the destinations that CSAir flies to, including its code, name, country, continent, timezone, longitude and latitude, population, region, and each of the cities that are accessible via a single non-stop flight from that destination
* Provide a graphical representation of CSAir's route map.

I chose to implement this in Go, as I had just been introducted to the language in my project with Robb Seaton (abdge) a few months earlier, building neptune. 

I used: http://www.gcmap.com/ to build a map, a JSON containing the various routes and information is included.

This was fairly straight forward, the fun part was using abstract{} in Golang and building a graphing library with a Dijkstra's Algorithm Implementation.

