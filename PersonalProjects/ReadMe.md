The first project is WebScraper. It takes a file with 1 million websites on it and uses worker pool with 100 workers, meaning it concurrently visits 100 websites at the same time, bringing down the execution time. 

The second project is TaxiDataAnalysis. It reads 18 different files, all containing New York Yellow Taxi Data from 2018 to 2020. The data can be found on Kaggle: https://www.kaggle.com/datasets/microize/newyork-yellow-taxi-trip-data-2020-2019. The data contains various information about the taxi rides like pickup date and time, dropoff date and time, number of passengers, total amount, etc. The code uploads the data to an existing SQL database. Code was also optimized to send 800 queries at once so that it runs faster
Both projects utilized channels and worker pool. 
