/**
 * Resume a string para o máximo de caracteres, mantendo as três primeiras partes e a última parte
 * @param text
 * @param maxLength
 */
export function resumeString(text: string, maxLength: number) {

    if (text?.length <= maxLength) {
        return text;
    }
    return text?.slice(0, maxLength - 3) + '...';
}

/**
 * Resume a string para o máximo de caracteres, mantendo as três primeiras partes e a última parte
 * @param text
 * @param maxLength
 */
export function resumeStringV2(text: string, maxLength: number) {
    const parts = text.split('/');
    if (parts.length <= 3) {
        return text;
    }

    // Juntar as três primeiras partes com "/" e depois adicionar "..." e a parte restante
    let result = parts.slice(0, 3).join('/') + '/...';

    // Limitar a parte final para ter exatamente maxLength caracteres
    const remainingLength = maxLength - result.length;
    result += parts.slice(3).join('/').slice(0, remainingLength);

    return result;
}

/**
 * Função para chamar um callback quando a tecla Enter for pressionada
 * Usado preferencialmente com onKeyDown
 * @param event
 * @param callback
 */
export function onEnter(event: any, callback: () => void) {
    if (event.key === 'Enter') {
        callback()
    }
}