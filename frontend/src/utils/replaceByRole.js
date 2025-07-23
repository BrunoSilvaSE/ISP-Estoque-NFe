import { jwtDecode } from '../src/utils/jwtDecode.js';
import { decodificarPermissoesStrToVet } from '../src/utils/roleCoderAndDecoder.js';

export async function replaceByRole(token) {
    try {
        let userRolesArray = [];
        if (token) {
            const decodedToken = jwtDecode(token);
            if (decodedToken && decodedToken.payload) {
                const roleString = decodedToken.payload.role;
                userRolesArray = await decodificarPermissoesStrToVet(roleString);
            }
        }

        const primaryRole = userRolesArray[0];

        switch (primaryRole) {
            case 'paciente':
                window.location.replace("/paciente");
                break;
            case 'admin':
                window.location.replace("/admin");
                break;
            case 'medico':
            case 'enfermeiro':
                window.location.replace("/main");
                break;
            case 'outros':
                window.location.replace("/main/ACS");
                break;
            default:
                window.location.replace("/login");
                break;
        }
    } catch (error) {
        alert("Ocorreu um erro ao processar as permissões do usuário.");
    }
}