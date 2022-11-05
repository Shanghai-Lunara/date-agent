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
    reset: $('#reset')[0],
    jobs: $('#body table tbody')[0],
    logs: $('#log')[0],
    modalModify: $('#modify')[0],
    // inputs: $('.body form input, .body form select, .test input'),
    // change: $('.change input')[0],
    resetValue: $('.reset input')[0],
    changeType: $('#changeType')[0],
    changeHours: $('#changeHours')[0],
    changeDays: $('#changeDays')[0],
    changeMins: $('#changeMins')[0],
    changeHoursBtn: $('#changeHoursBtn')[0],
    editing: undefined,
    taskId: 0,
    init: function () {
        app.getResponse();
        //Buttons
        app.getJobs();
        // app.new.on('click', app.newJob);
        app.close.on('click', app.closeModify);
        app.save.on('click', app.changeCommand);
        app.reset.on('click', app.resetCommand);
        app.changeHoursBtn.on('click', app.changeHoursCommand);
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
                  <td>${job.taskId}</td>
                  <td>${job.output}</td>
               </tr>`;
            app.jobs.innerHTML += template;
        });
    },
    openJob: function () {
        app.editing = this.id.split('_')[1];
        /*for(var i = 0; i < app.inputs.length; i++){
           app.inputs[i].value = app.data.rows[app.editing][app.inputs[i].getAttribute('name')] || "";
        }*/
        app.openModify();
    },
    changeCommand: function(){app.saveJob('date')},
    resetCommand: function() {app.saveJob(app.resetValue.value)},
    changeHoursCommand: function() {
        const type = app.changeType.value==='+' ? '+' : '-'
        const day = Number(app.changeDays.value)? Number(app.changeDays.value) : 0
        const hour = Number(app.changeHours.value)? Number(app.changeHours.value) : 0
        const minute = Number(app.changeMins.value)? Number(app.changeMins.value) : 0
        const changeTime = day*24*60 + hour*60 + minute
        const cmd = `date -d '${type}${changeTime}' minute +%Y-%m-%d %H:%M:%S`
        app.saveJob(encodeURIComponent(cmd))
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
    saveJob: function (request) {
        // editing[app.inputs[i].getAttribute('name')] = app.inputs[i].value;
        // let body = 'hostname=' + app.data.rows[app.editing]['hostname'] + '&command=' + app.inputs[i].value
        let body = 'command=' + request
        console.log('body2', body)
        app.ajax('post', '/changeTime', body)
        // app.data.rows[app.getJobIndex(app.editing)] = editing;
        // app.closeModify();
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
        xhr.send(body);

        xhr.onreadystatechange = function () {
            if (xhr.readyState == 4) { // 读取完成
                if (xhr.status == 200) {
                    console.log(JSON.parse(xhr.responseText));
                    const ret = JSON.parse(xhr.responseText).ret;
                    const tasks = JSON.parse(xhr.responseText).tasks;
                    if (url == '/getHub') {
                        let template = '';
                        Object.keys(tasks)?.forEach(function (key) {
                            let tmp = tasks[key]
                                app.taskId = tmp.id
                                template += `
                                <p>[TaskId:  ${tmp.id}] <label style="color: brown">Command:</label>     ${tmp.command} </p>
                                `;
                                app.logs.innerHTML = template;
                        })
                    }

                    data.rows = []
                    if(ret)
                    Object.keys(ret).forEach(function (key) {
                        let tmp = {
                            hostname: key,
                            taskId: ret[key]['task_id'],
                            output: ret[key]['output']
                        };
                        data.rows.push(tmp);
                    })
                    app.loadJobs();
                }
            }
        }
    },
    getResponse: function () {
        setInterval(function () {
            app.ajax('post', '/getHub')
        }, 200);
    },
    getDate: function () {
        var date = new Date(new Date().getTime());//如果date为13位不需要乘1000
        var Y = date.getFullYear() + '-';
        var M = (date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1) + '-';
        var D = (date.getDate() < 10 ? '0' + (date.getDate()) : date.getDate()) + ' ';
        var h = (date.getHours() < 10 ? '0' + date.getHours() : date.getHours()) + ':';
        var m = (date.getMinutes() < 10 ? '0' + date.getMinutes() : date.getMinutes()) + ':';
        var s = (date.getSeconds() < 10 ? '0' + date.getSeconds() : date.getSeconds());
        return Y + M + D + h + m + s;
    }

}

app.init();