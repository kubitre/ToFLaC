export function postRequest(request){
    fetch(request)
    .then((Response) => Response.json())
    .then((data) => data)
    .catch(err => console.log(err));
}

export function CreatePacket (API_URL, Body){
    let myHeaders = new Headers({
        "Content-Type": "application/json",
        "Accept": "application/json",
        "Access-Control-Allow-Origin": "*",
        "Origin": "*",
    });

    let myInit = { 
        method: "POST",
        headers: myHeaders,
        mode: "cors",
        cache: 'default',
        body: Body
    };

    let request = new Request(API_URL, myInit)
    return request
}