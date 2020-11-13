# Geocalc
![](https://gitlab.com/nicolasdscp/geocalc/badges/master/pipeline.svg)
![](https://gitlab.com/nicolasdscp/geocalc/badges/master/coverage.svg)
![](https://img.shields.io/badge/License-GPLv3-blue.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/govitia/geocalc)](https://goreportcard.com/report/github.com/govitia/geocalc)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/govitia/geocalc)](https://pkg.go.dev/github.com/govitia/geocalc)

Go minimalist library to compute earth points. Implementation of https://www.movable-type.co.uk/scripts/latlong.html.
Angle are in Rad and distance in meters.

## Features in v1.3.3

* Earth points in Rad/Degree
* Get distance between two Points
* Create Walk line
* Get bearing between two Points
* Rad/Degree conversion functions

## Getting start 

#### Creating Points
```go
// Point creation is in Rad
p1 := geocalc.NewPoint(math.Pi, 0)
p2 := geocalc.NewPoint(0, 0)
```
```go
// Get degree values for p1
latDegree, lonDegree := p1.Degree()
```

#### Bearing and distance
```go
// Get bearing in Rad between two Points
bearing := p1.Bearing(p2)
bearing := p2.Bearing(p1)
// Distance in meters
distance := geocalc.Distance(p1, p2)
```

#### Unit conversion
```go
rad := math.Pi
degree := geocalc.RadToDegree(rad) // 180
rad = geocalc.DegreeToRad(degree)  // math.Pi
```

#### Walking on the earth surface
Walk from a starting point to a final point with a bearing of 0 Rad and a distance of 1000m
```go
startPoint := geocalc.NewPoint(0, 0)
finalPoint := startPoint.Walk(0, 1000) // Return a new point 
```

#### Intermediate Points

`Intermediate` calculate an intermediate point at any fraction along the great circle path between
p1 and p2. Start path at p1, if fraction=0, returned Point is p1. If fraction=1,
returned Point is p2.

```go
p1 := geocalc.NewPoint(math.Pi, 0)
p2 := geocalc.NewPoint(0, 0)
fraction := 0.5 // We want a Point in the middle of the other two Points
p3 := geocalc.Intermediate(p1, p2, fraction)
```
