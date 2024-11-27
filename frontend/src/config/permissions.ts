export const ALL_ADMIN_PERMSSION = ['admin', 'super_admin', 'owner'];

export const ADMIN_PERMS = ['admin', 'super_admin'];
export const SUPER_ADMIN_PERMS = ['super_admin'];

export const hasAllAdminPermission = (role: string | null): boolean => {
  if (!role) return false; // Si le rôle est null, renvoie false
  return ALL_ADMIN_PERMSSION.includes(role); // Vérifie si le rôle est dans ALL_PERMS
};
