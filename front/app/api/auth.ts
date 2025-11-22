type Auth = {
    token: string
}

export async function login(email: string, password: string) {
    try {
        let res: Auth = await $fetch('api/login', {
            'method': 'POST',
            body: JSON.stringify({
                email,
                password
            })
        });
    } catch (err) {
        console.error("Login failed:", err)
    }
}