type Auth = {
    token: string
}

export async function login(email: string, password: string) {
    try {
        let res: Auth = await $fetch('http://localhost:8080/api/login', {
            method: 'POST',
            headers: {
                "Access-Control-Allow-Methods": 'GET,HEAD,PATCH,POST,DELETE',
                "Access-Control-Allow-Origin": '*',
                "Access-Control-Allow-Credentials": "true",
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                email,
                password
            }),
        });
        console.log(res);
    } catch (err) {
        console.error("Login failed:", err)
    }
}