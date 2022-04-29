const requestUrl = "localhost:8080/Klinefelter";
var input = document.getElementById("searchTxt");
function searchURL() {
    if (input.value.split(' ').length <= 3) {
        if (/^\d/.test(input.value)) {
            window.location.href = "http:localhost:8080/date/" + input.value.replace(/\s/g, '');
        } else {
            window.location.href = "http:localhost:8080/disease/" + input.value;
        }
    } else {
        if (/^\d/.test(input.value)) {
            var penyakit = input.value.split(" ").pop();
            var index = input.value.lastIndexOf(" ");
            input.value = input.value.substring(0, index); 
            window.location.href = "http:localhost:8080/dnd/" + penyakit + "/" + input.value.replace(/\s/g, '');
        } else {
            var index = input.value.lastIndexOf(" ");
            input.value = input.value.substring(0, index); 
            window.location.href = "http:localhost:8080/dnd/" + input.value.split(" ").pop() + "/" + input.value.replace(/\s/g, '');
        }
        
    }

}