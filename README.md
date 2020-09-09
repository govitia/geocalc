# Geocalc
![](https://gitlab.com/psns/geocalc/badges/develop/pipeline.svg)
![](https://gitlab.com/psns/geocalc/badges/develop/coverage.svg)
![](https://img.shields.io/badge/License-GPLv3-blue.svg)

Minimalist coordinate computation library

## API Reference

All these formulas are for calculations on the basis of a spherical earth
(ignoring ellipsoidal effects) – which is accurate enough for most purposes…
In fact, the earth is very slightly ellipsoidal; using a spherical model gives errors typically up to 0.3%.
More informations [here](https://www.movable-type.co.uk/scripts/latlong.html).

### Index

[Constants](#constants)
<br>[func Distance](#func-distancehttpsgitlabcompsnsgeocalc-blobdevelopgeocalcgol19)
<br>[func Intermediate](#func-intermediatehttpsgitlabcompsnsgeocalc-blobdevelopgeocalcgol33)
<br>[func DegreeToRad](#func-degreetoradhttpsgitlabcompsnsgeocalc-blobdevelopgeocalcgol38)
<br>[func RadToDegree](#func-radtodegreehttpsgitlabcompsnsgeocalc-blobdevelopgeocalcgol43)
<br>[type Point](#type-pointhttpsgitlabcompsnsgeocalc-blobdeveloppointgol6)
* [func NewPoint](#func-newpointhttpsgitlabcompsnsgeocalc-blobdeveloppointgol13)
* [func Bearing](#func-bearinghttpsgitlabcompsnsgeocalc-blobdeveloppointgol21)
* [func Walk](#func-walkhttpsgitlabcompsnsgeocalc-blobdeveloppointgol29)
* [func Degree](#func-degreehttpsgitlabcompsnsgeocalc-blobdeveloppointgol37)

### Constants
```go
const (
	EarthRadius        = 6371e3                    // EarthRadius in meters
	EarthCircumference = 2 * EarthRadius * math.Pi // EarthCircumference in meters
)
```

### func [Distance](https://gitlab.com/psns/geocalc/-/blob/develop/geocalc.go#L19)
```go
func Distance(p1, p2 Point) float64
```
`Distance` return the distance between two earth `Point` in meters

### func [Intermediate](https://gitlab.com/psns/geocalc/-/blob/develop/geocalc.go#L33)
```go
func Intermediate(p1, p2 Point, fraction float64) Point
```
`Intermediate` calculate an intermediate point at any fraction along the great circle path between
`p1` and `p2`. Start path at `p1`, if `fraction`=0, returned `Point` is `p1`. If `fraction`=1,
returned `Point` is `p2`. Number upper 1 can be used.

### func [DegreeToRad](https://gitlab.com/psns/geocalc/-/blob/develop/geocalc.go#L38)
```go
func DegreeToRad(deg float64) float64
```
DegreeToRad convert given arc in degree to rad

### func [RadToDegree](https://gitlab.com/psns/geocalc/-/blob/develop/geocalc.go#L43)
```go
func RadToDegree(rad float64) float64
```
RadToDegree convert given arc in rad to degree

### type [Point](https://gitlab.com/psns/geocalc/-/blob/develop/point.go#L6)
```go
type Point struct {
	Lat float64 // Latitude of the point in rad
	Lon float64 // Longitude of the point in rad
}
```
`Point` define an Earth point.

#### func [NewPoint](https://gitlab.com/psns/geocalc/-/blob/develop/point.go#L13)
```go
func NewPoint(lat, lon float64) Point 
```
`NewPoint` create and return an new `Point` from given coordinates, `lat` and `lon` in radians.

#### func [Bearing](https://gitlab.com/psns/geocalc/-/blob/develop/point.go#L21)
```go
func (p Point) Bearing(pt Point) float64
```
`Bearing` return the bearing in rad between current and given `Point`

#### func [Walk](https://gitlab.com/psns/geocalc/-/blob/develop/point.go#L29)
```go
func (p Point) Walk(bearing, dist float64) Point
```
`Walk` return destination `Point` in function of given `dist` and `bearing` from start p `Point`.
bearing in rad (clockwise from north) and dist in meters.

#### func [Degree](https://gitlab.com/psns/geocalc/-/blob/develop/point.go#L37)
```go
func (p Point) Degree() (float64, float64)
```
`Degree` return degree value of `Lat`, `Lon`