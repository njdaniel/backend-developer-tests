#!/bin/bash

echo "Starting endpoint tests*************"
echo "*****************"

echo "Testing GET /people  *****************"

echo "  testing for success 200"
get_allpeople_200=$(curl -s -o /dev/null -w "%{http_code}" localhost:9000/people)

if [[ "${get_allpeople_200}" =~ ^200 ]]; then
    echo "SUCCESS: Got $get_allpeople_200"
else
    echo "FAILED wanted 200, got $get_allpeople_200"
fi
echo ""



echo "Testing GET /people/{ID}  *****************"
echo "  testing for success 200"


get_peopleid_200=$(curl -s -o /dev/null -w "%{http_code}" localhost:9000/people/135af595-aa86-4bb5-a8f7-df17e6148e63)
if [[ "${get_peopleid_200}" =~ ^200 ]]; then
    echo "SUCCESS: Got $get_peopleid_200"
else
    echo "FAILED wanted 200, got $get_peopleid_200"
fi


get_peopleid_404=$(curl -s -o /dev/null -w "%{http_code}" localhost:9000/people/135af595-aa86-4bb5-a8f7-di17e6198e69)
echo "  testing for not found 404"

if [[ "${get_peopleid_404}" =~ ^404 ]]; then
    echo "SUCCESS: Got $get_peopleid_404"
else
    echo "FAILED wanted 404, got $get_peopleid_404"
fi
echo ""


echo "Testing GET /people?first_name=:first_name&last_name=:last_name  *****************"
echo "  testing for success 200"
get_peoplename_200=$(curl -s -o /dev/null -w "%{http_code}" 'localhost:9000/people?first_name=John&last_name=Joe')
if [[ "${get_peoplename_200}" =~ ^200 ]]; then
    echo "SUCCESS: Got $get_peoplename_200"
else
    echo "FAILED wanted 200, got $get_peoplename_200"
fi
#echo "  testing for result"
#get_people_result=$(curl -s  -X GET 'localhost:9000/people?first_name=John&last_name=Doe')
#echo "result: $get_people_result"
#if [[ "${get_people_empty}" =~ ^\[()?!^$)([^\s])]  ]]; then
#    echo "SUCCESS"
#else
#   echo "FAILED wanted [], got $get_people_empty" 
#fi
echo "  testing for empty result 200"
get_peoplename_200=$(curl -s -o /dev/null -w "%{http_code}" 'localhost:9000/people?first_name=Donny&last_name=Zany')
if [[ "${get_peoplename_200}" =~ ^200 ]]; then
    echo "SUCCESS: Got $get_peoplename_200"
else
    echo "FAILED wanted 200, got $get_peoplename_200"
fi
#echo "  testing for empty array"
#get_people_empty=$(curl -s -X GET 'localhost:9000/people?first_name=Donny&last_name=Zany')
#echo " empty: $get_people_empty"
#if [[ "${get_people_empty}" =~ ^\[]  ]]; then
#    echo "SUCCESS"
#else
#   echo "FAILED wanted [], got $get_people_empty" 
echo ""

echo "Testing GET /people?phone_number=\+1 \(800\)\ 555\-1313 *****************"
echo "  testing for success 200"
get_peoplephone_200=$(curl -s -o /dev/null -w "%{http_code}" localhost:9000/people?phone_number=\+1 \(800\)\ 555\-1313)
if [[ "${get_peoplephone_200}" =~ ^200 ]]; then
    echo "SUCCESS: Got 200"
else
    echo "FAILED wanted 200, got $get_peoplephone_200"
fi

echo "  testing for empty result 200"
get_peoplephone_200=$(curl -s -o /dev/null -w "%{http_code}" localhost:9000/people?phone_number=\+1 \(800\)\ 555\-1313)
if [[ "${get_peoplephone_200}" =~ ^200 ]]; then
    echo "SUCCESS: Got 200"
else
    echo "FAILED wanted 200, got $get_peoplephone_200"
fi
echo ""
