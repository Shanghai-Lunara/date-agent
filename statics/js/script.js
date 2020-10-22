/*----------------| Database |------------------*/
var data = {
    id: 2,
    rows: [
        /*{
           key: 0,
           name: 'Task 1',
           path: '/some/path',
           startDate: '2017-02-15',
           startTime: '00:00:00',
           endDate: '',
           endTime: '',
           repeatEvery: undefined,
           interval: 'Daily',
           days: '3,5',
           weekNums: undefined
        },*/
    ]
}

/*----------------| Application |------------------*/

//Fake some JQuery
var $ = function (str) {
    var ele = document.querySelectorAll(str);
    for (var i = 0; i < ele.length; i++) {
        ele[i].on = function (event, func) {
            this.addEventListener(event, func);
        }
    }
    return ele;
};

//Fake a Framework
var app = {
    data: data,
    new: $('#new')[0],
    close: $('#close')[0],
    del: $('#del')[0],
    save: $('#save')[0],
    jobs: $('#body table tbody')[0],
    logs: $('#log')[0],
    modalModify: $('#modify')[0],
    inputs: $('.body form input, .body form select'),
    editing: undefined,
    taskId: 0,
    init: function () {
        app.getResponse();
        //Buttons
        app.getJobs();
        // app.new.on('click', app.newJob);
        app.close.on('click', app.closeModify);
        app.save.on('click', app.saveJob);
        // app.del.on('click', app.delJob);

        app.closeModify();
    },
    openModify: function () {
        app.modalModify.classList.remove('close');
    },
    closeModify: function () {
        app.loadJobs();
        app.modalModify.classList.add('close');
    },
    loadJobs: function () {
        app.jobs.innerHTML = "";
        this.data.rows.forEach(function (job, index) {
            var jobID = 'i_' + index;
            var template = `
               <tr class='jobrow' id='${jobID}'>
                  <td>${job.hostname}</td>
                  <td>${job.status}</td>
               </tr>`;
            app.jobs.innerHTML += template;
        });
        const rows = $('.jobrow');
        const len = rows.length || 1;
        for (let i = 0; i < rows.length; i++) {
            rows[i].on('click', app.openJob);
        }
    },
    openJob: function () {
        app.editing = this.id.split('_')[1];
        /*for(var i = 0; i < app.inputs.length; i++){
           app.inputs[i].value = app.data.rows[app.editing][app.inputs[i].getAttribute('name')] || "";
        }*/
        app.openModify();
    },
    /*newJob: function(){
       for(var i = 0; i < app.inputs.length; i++){
          app.inputs[i].value = "";
       }
       $('input[name=key]').value = app.data.id;
       app.editing = app.data.id;
       app.data.id += 1;
       app.openModify();
    },*/
    saveJob: function () {
        for (let i = 0; i < app.inputs.length; i++) {
            // editing[app.inputs[i].getAttribute('name')] = app.inputs[i].value;
            let body = 'hostname=' + app.data.rows[app.editing]['hostname'] + '&command=' + app.inputs[i].value
            app.ajax('post', '/changeTime', body)
            console.log(app.inputs[i].value);
            console.log(app.data.rows[app.editing]['hostname']);
        }
        // app.data.rows[app.getJobIndex(app.editing)] = editing;
        app.closeModify();
    },
    /*delJob: function(){
       app.data.rows.splice(app.editing, 1);
       app.closeModify();
    },*/
    /*getJobIndex: function(key){
       let i;
       for(i = 0; i < app.data.rows.length; i++){
          if(app.data.rows[i].key == Number(key))
              break;
       }
       return i;
    },*/
    getJobs: function () {
        app.ajax();
    },
    ajax: function (method = 'get', url = '/getJobs', body = null) {
        const xhr = new XMLHttpRequest();
        xhr.open(method, url, true);
        if (method == 'post') {
            xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
        }
        console.log(body);
        xhr.send(body);

        app.loadJobs();
        xhr.onreadystatechange = function () {
            if (xhr.readyState == 4) { // 读取完成
                if (xhr.status == 200) {
                    console.log(url)
                    console.log(xhr.responseText);
                    if (url != '/getHub') {
                        data.rows = []
                        Object.keys(JSON.parse(xhr.responseText)).forEach(function (key) {
                            console.log(key, JSON.parse(xhr.responseText)[key]);
                            data.rows.push(JSON.parse(xhr.responseText)[key]);
                        })
                        console.log(data.rows);
                        app.loadJobs();
                    } else {
                        Object.keys(JSON.parse(xhr.responseText)).forEach(function (key) {
                            console.log(key, JSON.parse(xhr.responseText)[key]);
                            let tmp = JSON.parse(xhr.responseText)[key]
                            if (tmp.id > app.taskId) {
                                app.taskId = tmp.id
                                let template = '';

                                Object.keys(tmp.result).forEach(function (key) {
                                    template += `
                                    <p>[${app.getDate()}]  ${tmp.result[key]}</p>
                                    `;
                                    console.log(template);
                                })
                                console.log(template);
                                app.logs.innerHTML += template;
                            }
                        })
                    }
                }
            }
        }
    },
    getResponse: function () {
        setInterval(function () {
            app.ajax('post', '/getHub')
        }, 3000);
    },
    getDate: function () {
        var date = new Date(new Date().getTime());//如果date为13位不需要乘1000
        var Y = date.getFullYear() + '-';
        var M = (date.getMonth()+1 < 10 ? '0'+(date.getMonth()+1) : date.getMonth()+1) + '-';
        var D = (date.getDate() < 10 ? '0' + (date.getDate()) : date.getDate()) + ' ';
        var h = (date.getHours() < 10 ? '0' + date.getHours() : date.getHours()) + ':';
        var m = (date.getMinutes() <10 ? '0' + date.getMinutes() : date.getMinutes()) + ':';
        var s = (date.getSeconds() <10 ? '0' + date.getSeconds() : date.getSeconds());
        return Y+M+D+h+m+s;
    }

}

app.init();