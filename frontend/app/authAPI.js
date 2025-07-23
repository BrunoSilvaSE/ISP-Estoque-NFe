import "/environment/environment.js";
import { replaceByRole } from "/src/utils/replaceByRole";


export async function loginRequisicao(cpf, psw) {
    const credenciais = { cpf, psw };

    const response = await fetch(window.API_ENDERECO + "auth/login", {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(credenciais)
    });

    if (!response.ok) {
        const errorData = await response.json().catch(() => ({ message: `Erro HTTP: ${response.status}` }));
        throw new Error(errorData.message);
    }

    const tokenObj = await response.json();
    
    localStorage.setItem("token", tokenObj.token);

    await replaceByRole(tokenObj.token);
}
