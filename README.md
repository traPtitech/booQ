# booQ
[![CircleCI](https://circleci.com/gh/traPtitech/booQ/tree/master.svg?style=shield)](https://circleci.com/gh/traPtitech/booQ/tree/master)

management tool for equipment and book rental

## Development environment
### Setup with docker and docker-compose
#### First Up (or entirely rebuild)
```
$ docker-compose up --build
```

Now you can access to `http://localhost:3000` for booQ

And you can access booQ MariaDB by executing commands
`docker-compose exec db bash` and `mysql -uroot -ppassword -Dbooq` 

You can also execute build shell script
```
$ sh scripts/build.sh
```
And start development
```
$ sh scripts/up.sh
``` 

#### test
You can test this project
```
$ docker-compose -f docker/test/docker-compose.yml up --abort-on-container-exit
```

Or you can also execute test shell script
```
$ sh scripts/test.sh
```

#### Rebuild
`docker-compose up --no-deps --build`

#### Destroy Containers and Volumes
`docker-compose down -v`

## Contribution
The task list is in [issue](https://github.com/traPtitech/booQ/issues)

1. Fork it ( https://github.com/traPtitech/booQ )
2. Create your feature branch (git checkout -b issue-{{number}}) e.g. `issue-1`
3. Commit your changes (git commit -m 'Add some feature')
4. Push to the branch (git push origin issue-{{number}})
5. Create new Pull Request
