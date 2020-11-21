function check(typ){}
function getid(elid){var f=document.getElementById(elid);return f;}
function gf(formname,fieldname){var f=eval("document."+formname+"."+fieldname);return f.value;}
function showresult(id,val){var rs=getid('r'+id);rs.value=val;}
$(document).ready(function()
{$("textarea.auto_selectemb").mouseover(function()
{$(this).select();});$("#embedisp").hide();});function embedshowhide(obj)
{if(obj=='expand')
{$("#embedisp").show();}
else if(obj=='minimize')
{$("#embedisp").hide();}}