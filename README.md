
# url-shortener

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)

---

A service for shortening links using the Go programming language and packages such as GORM to interact with the PostgreSQL database and Fiber to create a web interface.

## Deployment

To deploy this project, run the following command

```bash
make build && ./bin/url-shortener
```

It starts the gofiber server on the port specified in the `.env` file

## API Reference

#### Using a redirect to another site with a short name

```http
  GET /i/${shortname}
```

| Parameter   | Type     | Description                       |
| :-----------| :------- | :-------------------------------- |
| `shortname` | `string` | **Required**. TThe short name is the unique identifier assigned to the shortened link. When this endpoint is accessed with the appropriate short name, it redirects the user to the original URL associated with that short name |

#### Creating a new shortened link

```http
POST /api/v1/shortlinks/create
```

| Parameter   | Type     | Description                                                                                                                                                    |
|:------------|:---------|:---------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `api_key`   | `string` | **Required**. The API_KEY is a unique identifier. It is necessary for authentication and access control purposes                                               |
| `shortname` | `string` | **Required** The SHORTNAME is a user-defined alias or identifier for the shortened link. It helps users remember and identify the link easily                  |
| `url`       | `string` | **Required** The URL is the original long URL that needs to be shortened. It is the destination address that the shortened link will redirect to when accessed |

## License

This project is distributed under the terms of the [MIT](https://choosealicense.com/licenses/mit/) license.
