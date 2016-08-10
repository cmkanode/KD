I built most of the webserver using the following website:
http://stevenwhite.com/building-a-rest-service-with-golang-2/

Did a lot of searching on Google to help find solutions.  
I ran into an issue with getting the document from MongoDB to map to my model 
in Go.  Case was the issue.
--------------------------------------------------------------------------------
Issues:
• Retrieving a Title by TitleId is not working.
    - I don't really need it.  If I'm return the search results to the client,
        then the client can pull the individual records that the user may want
        to view.
• Ran out of time, so I didn't get a chance to build out the models for the children
    - Awards, Genres, OtherNames, Participants, StoreLines
    - Was disappointed that I didn't have time to play with React.js, though I did
        have a lot of fun working with Go.

--------------------------------------------------------------------------------
Possible Environmental Issues:
• My Win10 tablet is 32bit.  I was unable to set up the same environment as on my 
    desktop.  Two issues did come up for that environment:  
    - fmt.Fprint did not work correctly.  It ignored "%s"
    - Unable to install gopkg.in/mgo.v2/bson.  Based on the error message, there
        was an int overflow.
• Work PC.  I was going to take a little time during my breaks to add more to my
    Go code, but upon intial install, my code does not compile.  The error points
    to a specific line, but I don't have enough details to determine why it's 
    breaking there for this environment, but compiles fine on my home PC.
--------------------------------------------------------------------------------
Post-Mortem:
• I spent too much time getting the environment set up on both my main PC and my tablet.
    - Had to add an environmental variable, GOPATH, and had to close and re-open
        the Command window mutiple times after making changes.
• I made assumptions.
    - Biggest assumption was thinking that the Collection name was lower case: "titles".
        This caused me to spend too much time trying to trouble-shoot my connection
        to the database.
• I decided that it was worth the risk to work in Go
    - The side effect was that I burned through my time, and did not have enough
        to work on the GUI other than throw some basic work together.
