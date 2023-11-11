
# GROUPIE-TRACKER-search-bar
* `ayessenb` 
* `roshakba` 



#### Groupie-tracker-search-bar consists on making your site :

    more appealing, interactive and intuitive.
    more user friendly.
    give more feedback to the user.
    has JSON file which contains the data that need to be displayed
    Additionally for the search bar we have added the features which, allow for user to enter the string
    and the search bar will be suggesting any string based on the input from the JSON file, which could be:
   
    * artist/band name
    * members
    * locations
    * first album date
    * creation date

    If any of the above parameters contains the entered substring to the search bar, 
    that will be displaying dynamically the list of all possible suggestions for the user 
    and also the page dynamically will be displaying only those bands and/or artists 
    on the main page that are in the list of suggestions.


#### Description:

* We used text/template pakage to recieve and send GET requests.
* GET /: Sends HTML response, the main page.
* and using the JSON file which contains the data that need to be displayed to pages
* Linked external css file to make good design 
* For realising the dynamnic chnages of the bands and list of suggestions,
 JS was used for the implementation on the front ent.



## Usage/Examples
Cloning storage to your host
```CMD/Terminal 
git clone git@github.com:compygirl/groupie-tracker-search-bar.git
```
Go to the downloaded repository:

```CMD/Terminal 
cd groupie-tracker-search-bar
```
Run a Server:
```CMD/Terminal 
go run main.go
```

Follow the link on the terminal:
```CMD/Terminal 
Starting server got testing... http://127.0.0.1:8080 
```

you can play with the page by choosing the music band's image and so on.
As well as seaching the strings



## HTTP status codes
* OK (200), if everything went without errors.
* Not Found, if the wrong URL was provided.
* Bad Request, for incorrect requests. for exaple, the id is out of range.
* Internal Server Error, for unhandled errors.
* If string is not a substring of any items which were mentioned in the first paragraph, 
then the page will be displaying "NOT FOUND"



## Feedback

If you liked our project, we would be grateful if you could add `Star` to the repository.

Alem Student
11.08.2023.