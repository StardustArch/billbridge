


export const saveAuthData = (token, username) => {
    localStorage.setItem('authToken', token);
    localStorage.setItem('authUser', JSON.stringify({ username }));
  };
  
  
  // Obtém o token
  export const getToken = () => {
    return localStorage.getItem('authToken');
  };
  
  // Obtém o nome de usuário
  export const getUsername = () => {
    const userJson = localStorage.getItem('authUser');
    if (!userJson) return null;
    try {
      return JSON.parse(userJson).username;
    } catch {
      return null;
    }
  };
  
  // Remove tudo (logout)
  export const clearAuthData = () => {
    localStorage.removeItem('authToken');
    localStorage.removeItem('authUser');
  };
  
  // Verifica se está autenticado
  export const isAuthenticated = () => {
      return !!getToken();
    };
  
  export default {
    saveAuthData,
    getToken,
    getUsername,
    clearAuthData,
    isAuthenticated
  };
  