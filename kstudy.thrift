namespace go api

struct Request {
    1: string message1
    2: string message2
}

struct Response {
    1: string message
}

service KStudy {
    Response Concat(1: Request req)
}