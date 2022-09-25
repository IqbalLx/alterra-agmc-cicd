# Go + Docker Deploy

## Stacks used
- Database: planetscale -> serverless MySQL Database
- Server: heroku -> actually it's serverless too, so we don't need to provisioned any server manually

## Steps
- Configure Dockerfile
  
  Because of heroku is serverless, it needs some spesific Dockerfile configuration to match heroku specs. We can't expose PORT because heroku assigned our app port automatically, this also require us to make Go read port from env instead of hardcoded value.

- Setup free [planetscale](https://planetscale.com/) database
  ![planetscale](docs/Screen%20Shot%202022-09-23%20at%2021.39.42.png "planetscale")

- Setup new heroku app
  ![heroku](docs/Screen%20Shot%202022-09-23%20at%2021.38.16.png "heroku")

- Configure env variable inside heroku
  ![env](docs/Screen%20Shot%202022-09-23%20at%2021.38.52.png "env")

- Build and push container using Heroku CLI
  ![](docs/Screen%20Shot%202022-09-23%20at%2021.40.07.png)
  ![](docs/Screen%20Shot%202022-09-23%20at%2021.40.16.png)


The endpoint avaliable here: https://alterra-agmc-iqbal.herokuapp.com/v1
The API Docs can be found here: https://documenter.getpostman.com/view/23104123/2s7ZLkoWdz

Cheers