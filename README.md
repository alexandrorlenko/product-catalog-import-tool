# The script for a product feed validation, transformation and uploading to the predefined system


## Preamble
We are Setronica company with the big experience to build B2B and B2C integrations to exchange product catalog information between PIMs, Marketplaces, Storefronts, ERPs, services, and platforms. We know that it is quite difficult to integrate with a new service or a platform and we defined two the most important reasons:
    
1.  Each service has its native format and API format. It takes time to meet and configure it properly. And even if a service provides a lot of integration flexibility then it still requires to meet with this flexibility and again to configure it properly;
2.  Each seller has its format of data with which they are used to working or just because their PIM system supports it.

This problem can be solved in three ways: on your side, on the system side, and by finding an external service that helps with the integration to the system. Each of them has its pros and cons. It is up to you to choose the best one for you. You may be lucky to have the integration from the system as out of the box for you. But all others who find it painful are welcome. This solution was built for you.

We did it in all three ways and our conclusion that the seller side is the most perspective.
So we tried to implement something simple to have the ability to integrate with the service quickly without diving deeply inside of it. This script can be run locally which allows for you to get proof of the working process as soon as it is configured without any infrastructure challenges. Later it can be deployed somewhere to provide automatization and autonomy.

We implemented four steps where each of them can be switched on/off if it isn’t applicable for your case:

1.  Map data from your field names to the system field names; 
2.  Validate data based on the system rules;
3.  Transform data into the system’s native format;
4.  Send data to the system.

The script  is supporting:

*  the integration with the following systems: Tradeshift.
*  the following formats of incoming data: CSV.
*  the following formats of outcoming data: CSV, EHF (in the nearest future)

## Files structure
We have four main folders:

1.  Mapping - this folder contains the file with a mapping of a seller’s fields and the  system’s fields. Mapping file should be in '.yaml' format and must contain required fields:
    *  ID
    *  Category
2.  Ontology - this folder contains the file with all rules that the system requests from their sellers. Each category has its list of attributes that have to be filled in a way as the system expects it.
   Ontology file should be in csv format with ',' separator.
3.  Source - this folder contains incoming files from a seller with the following subfolders
    *  ‘inprogress’ folder to move file here as soon as processing is started;
    *  ‘processed’ folders to move files when the processing is finished successfully and the source file has all required attributes to proceed with the next step.
4.  Result - this folder contains results of validation and transformation with the following subfolders:
    *  ‘report’ - while a file has any missing attributes then the script will keep it in this folder with all recommendations to fix; 
    *  ‘sent’ - If the file is correct and sent to the system, then you can find it in this folder.

## Configuration file
This [./service.yaml](service.yaml)  file contains settings to establish an API connection with the system.

The easiest way to get started working with the Tradeshift API is to create OAuth credentials by activating the API Access to Own Account app ([click here...]( https://sandbox.tradeshift.com/#/apps/Tradeshift.AppStore/apps/Tradeshift.APIAccessToOwnAccount)). 
The app will display your credentials. Just copy these values and paste them into the configuration file:

*  base_url
*  consumer_key
*  consumer_secret
*  token
*  token_secret
*  tenant_id

## Build
You need to run the following command to build the script and initialise default folders:


    ./install.sh

[Click to see more...](./INSTALL.md)

## How does it work?
Let’s just take a file in the supported format. If it isn’t in the supported format that some of them are compatible and can be converted like XLSX file can be saved as CSV.

1.  Place the script on an infrastructure and build it
2.  Place the system’s credentials into the configuration file to make a connection.
3.  Configure the mapping file to set your columns names on the right side of it.
4.  Place our file into the ‘source’ folder

        scp ./products.csv <user>@<host>:./data/source/

5.  Run the script 

        ssh <user>@<host>
        ./ts

6.  Download a report file from the ‘result/report’ folder and look at recommendation  messages inside if a required attribute is missed in the source file
   
        scp <user>@<host>:./data/result/report/products-failures.csv ./

7.  An operator can correct all problems and provide missed information by following instructions in the report file
8.  Put the corrected report file into the ‘source’ folder again
        
        scp ./products-failures.csv <user>@<host>:./data/source/

9.  Run the script
        
        ./ts

10.  Repeat steps (4)-(7) if the script is still reporting problems
11.  If everything is fine your file has been transformed and sent to the system.
You can see sending report here:
    
        ./data/result/report/produts_tradeshift-import-results.txt
   
12.  Celebrate!
