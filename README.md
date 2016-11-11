# Overwatch Hero API

The unofficial Overwatch Hero API allows developers to get data for every hero in the Overwatch universe as well as game types for competitive play.  

## Routes:

The base URL is (https://overwatch-hero-api.herokuapp.com/api/v1/)

To get a list of all heroes, send a get request to `https://overwatch-hero-api.herokuapp.com/api/v1/heroes`

For a single hero, the URL is `https://overwatch-hero-api.herokuapp.com/api/v1/heroes/:id`

Similarly, gametypes are accessed at `https://overwatch-hero-api.herokuapp.com/api/v1/gametypes`

and `https://overwatch-hero-api.herokuapp.com/api/v1/gametypes/:id`

#### [Overwatch Team Grader Deployed Site](https://isaacmillercodes.github.io/overwatch-team-grader/)

#### [Overwatch Team Grader Repo](https://github.com/isaacmillercodes/overwatch-team-grader)

**Overwatch Hero API was built with:**

* Go
* Gin
* Gorp
* PostgreSQL
