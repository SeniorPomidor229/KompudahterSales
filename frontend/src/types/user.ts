export interface User {
    _id: string;
    firstName: string;
    lastName: string;
    email: string;
    password: string;
    is_admin: boolean;
}
  
export interface UserDto {
    username: string;
    password: string;
}