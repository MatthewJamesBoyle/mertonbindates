# Merton Bin Collection dates
This app scrapes https://services.merton.gov.uk to provide the next bin collection dates. The intention was to wrap an Alexa skill around this. This can be found here: https://github.com/MatthewJamesBoyle/alexabindates

To get this working you will need to create a .env file in the route of the project with the following values:
```
URL=https://services.merton.gov.uk/tkflow_ppweb/flow.aspx?f=CollectionDay.kdt
HOUSE_NUMBER=
POSTCODE=
```
Please note if you do not provide a postcode and house number in the London borough of Merton, this application will not work for you.