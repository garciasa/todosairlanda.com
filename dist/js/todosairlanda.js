//auto expand textarea
function adjust_textarea(h) {
    h.style.height = "20px";
    h.style.height = (h.scrollHeight) + "px";
}
function init() {
    AOS.init();
    let anchorlinks = document.querySelectorAll('a[href^="#"]')
    for (let item of anchorlinks) {
        item.addEventListener('click', (e) => {
            let hashval = item.getAttribute('href')
            let target = document.querySelector(hashval)
            target.scrollIntoView({
                behavior: 'smooth',
                block: 'start'
            })
            history.pushState(null, null, hashval)
            e.preventDefault()
        })
    }
}
function contacta(e) {
    e.preventDefault();
    if (document.getElementsByName("nombre")[0].value === '' &&
        document.getElementsByName("email")[0].value === '') {
        document.getElementsByClassName("message")[0].innerHTML = "Nombre e Email son obligatorios"
        document.getElementsByClassName("message")[0].style.backgroundColor = "red";
        document.getElementsByClassName("message")[0].style.display = "block";
        return
    }
    var form = document.querySelector("form");
    var data = {}
    var formdata = new FormData(form)
    for (let tuple of formdata.entries()) data[tuple[0]] = tuple[1];
    var request = new XMLHttpRequest();
    request.onreadystatechange = function () {
        if (this.readyState == 4 && this.status == 200) {
            var jsonresp = JSON.parse(this.responseText);
            document.getElementsByClassName("message")[0].innerHTML = jsonresp.Message;
            document.getElementsByClassName("message")[0].style.backgroundColor = "green";
            document.getElementsByClassName("message")[0].style.display = "block";
            document.getElementById("formContacta").reset();
        }
    };
    request.open("POST", "/contacta", true);
    request.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
    request.send(JSON.stringify(data));
}
