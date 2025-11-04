export function generatePassword(length: number = 12): string {
    const lowercase = "abcdefghijklmnopqrstuvwxyz";
    const uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
    const numbers = "0123456789";
    const special = "!@#$%^&*()-_=+[]{}|;:,.<>?";

    const allChars = lowercase + uppercase + numbers + special;

    // Ensure at least one of each type
    const getRandom = (chars: string) =>
        chars[Math.floor(crypto.getRandomValues(new Uint32Array(1))[0] / (2 ** 32) * chars.length)];

    let password = [
        getRandom(lowercase),
        getRandom(uppercase),
        getRandom(numbers),
        getRandom(special),
    ];

    // Fill the rest randomly
    for (let i = password.length; i < length; i++) {
        password.push(getRandom(allChars));
    }

    // Shuffle the password to avoid predictable order
    password = password.sort(() => Math.random() - 0.5);

    return password.join("");
}



export function getRandomValue<T extends Record<string, any>>(obj: T): T[keyof T] {
    const values = Object.values(obj);
    const randomIndex = Math.floor(Math.random() * values.length);
    return values[randomIndex];
}
