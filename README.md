# SUT COURSE API (Go version) 📚

This project scrapes course data from Suranaree Uni of Tech's Reg. [Typescript ver.](https://github.com/pandakn/sut-course-api)

## Getting started 🚀

Clone this repository

```zsh
git clone https://github.com/pandakn/go-sut-course-api.git

cd go-sut-course-api
```

### Start Project

```zsh
go run main.go
```

## API Reference

#### Get Courses

```http
  POST /api/v1/courses
```

| Key                   | Type     | Description                                                                                                                                                                                                                                                     | Example           |
| :-------------------- | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | :---------------- |
| `acadYear`            | `string` | Academic year (e.g. 2565)                                                                                                                                                                                                                                       | `"2565"`          |
| `semester`            | `int`    | Semester (e.g. 3)                                                                                                                                                                                                                                               | `3`               |
| `courseCode`          | `string` | Course code pattern (e.g. "ist30 1105")                                                                                                                                                                                                                         | `"ist30 1105"`    |
| `courseName`          | `string` | Course name pattern (e.g. "english\*")                                                                                                                                                                                                                          | `"english*"`      |
| `maxRow`              | `int`    | Max rows in response (default is 50)                                                                                                                                                                                                                            | `25`              |
| `isFilter (optional)` | `bool`   | Filter by day and times (`true`) or no filter (`false`)                                                                                                                                                                                                         | `true` or `false` |
| `day (optional)`      | `string` | Weekdays for filtering courses (e.g. "monday") (e.g. "monday"). Use the format "sunday", "monday", ..., "saturday". Required if `isFilter` is `true`.                                                                                                           | `"monday"`        |
| `timeFrom (optional)` | `string` | Starting time for filtering courses (e.g. "08:00")                                                                                                                                                                                                              | `"08:00"`         |
| `timeTo (optional)`   | `string` | Ending time for filtering courses (e.g. "12:00")                                                                                                                                                                                                                | `"12:00"`         |
| `faculty (optional)`  | `string` | Faculty to filter the courses. Use the following values: `"ALL"`, `"SCIENCE"`, `"SOCIAL_TECHNOLOGY"`, `"AGRICULTURAL_TECHNOLOGY"`, `"MEDICINE"`, `"ENGINEERING"`, `"NURSING"`, `"DENTISTRY"`, `"PUBLIC_HEALTH"`, `"DIGITAL_ARTS_AND_SCIENCE"`. Default is `ALL` | `"SCIENCE"`       |

> **Note** : Time range : 08:00 - 22:00

> **Note** : Either one of courseCode or courseName can be entered.

> **Warning** : If courseCode and courseName are not specified, scraping all the data will take a very long time.

### Examples

#### Retrieve Course Data with No Filtering

use the following example:

```json
// body request
{
  "acadYear": "2566",
  "semester": 2,
  "courseCode": "523332",
  "courseName": "",
  "maxRow": 50
}
```

#### Retrieve Course Data with Filtering

```json
// body request
{
  "acadYear": "2566",
  "semester": 2,
  "courseCode": "523332",
  "courseName": "",
  "maxRow": 50,
  "isFilter": true,
  "day": "monday",
  "timeFrom": "10:00",
  "timeTo": "12:00"
}
```

## Usage/Examples JSON

```json
{
  "year": "2/2566",
  "faculty": "ALL",
  "courses": [
    {
      "courseCode": "523332",
      "version": "2",
      "courseName": {
        "en": "SOFTWARE ENGINEERING",
        "th": "วิศวกรรมซอฟต์แวร์"
      },
      "credit": "4 (3-3-9)",
      "degree": "ปริญญาตรี",
      "department": "วิศวกรรมคอมพิวเตอร์",
      "faculty": "สำนักวิชาวิศวกรรมศาสตร์",
      "courseStatus": "ใช้งาน",
      "courseCondition": ["523331"],
      "continueCourse": ["523435"],
      "equivalentCourse": null,
      "sectionsCount": 2,
      "sections": [
        {
          "id": "0baeef8b-98f8-4fdf-a16a-ef8a0c922d66",
          "url": "http://reg.sut.ac.th/registrar/class_info_2.asp?backto=home&option=0&courseid=1009172&coursecode=523332&acadyear=2566&semester=2&avs882850039=3",
          "section": "1",
          "status": "เปิดลงปกติ สามารถลงทะเบียนผ่าน WEB ได้",
          "note": "สำหรับหลักสูตรปรับปรุง พ.ศ. 2560",
          "professors": [
            "อาจารย์ ดร.คมศัลล์ ศรีวิสุทธิ์",
            "นายธนพล คงเจริญสุข",
            "นายสิทธิชัย สิริฤทธิกุลชัย",
            "นายตะวัน คำอาจ"
          ],
          "language": "TH",
          "seat": {
            "totalSeat": "45",
            "registered": "45",
            "remain": "0"
          },
          "classSchedule": [
            {
              "day": "Tu",
              "times": "09:00-12:00",
              "room": "B1139"
            },
            {
              "day": "Th",
              "times": "09:00-12:00",
              "room": "F11-422.Software"
            }
          ],
          "exams": {
            "midterm": {
              "date": "25",
              "month": "Dec",
              "times": "12:00-14:00",
              "year": "2566",
              "room": "อาคารBห้องB1115(สอบตามตารางมหาวิทยาลัย)25ธ.ค.2566"
            },
            "final": {
              "date": "14",
              "month": "Feb",
              "times": "13:00-16:00",
              "year": "2567",
              "room": "อาคารBห้องN(สอบตามตารางมหาวิทยาลัย)14ก.พ.2567"
            }
          }
        },
        {
          "id": "f8eedbc6-7c1b-4e3e-9015-348c4657802d",
          "url": "http://reg.sut.ac.th/registrar/class_info_2.asp?backto=home&option=0&courseid=1009172&coursecode=523332&acadyear=2566&semester=2&avs882850039=4",
          "section": "2",
          "status": "เปิดลงปกติ สามารถลงทะเบียนผ่าน WEB ได้",
          "note": "สำหรับหลักสูตรปรับปรุง พ.ศ. 2560",
          "professors": [
            "อาจารย์ ดร.คมศัลล์ ศรีวิสุทธิ์",
            "นายธนพล คงเจริญสุข",
            "นายสิทธิชัย สิริฤทธิกุลชัย",
            "นายตะวัน คำอาจ"
          ],
          "language": "TH",
          "seat": {
            "totalSeat": "40",
            "registered": "38",
            "remain": "2"
          },
          "classSchedule": [
            {
              "day": "Tu",
              "times": "09:00-12:00",
              "room": "B1139"
            },
            {
              "day": "Th",
              "times": "13:00-16:00",
              "room": "F11-422.Software"
            }
          ],
          "exams": {
            "midterm": {
              "date": "25",
              "month": "Dec",
              "times": "12:00-14:00",
              "year": "2566",
              "room": "อาคารBห้องB1115(สอบตามตารางมหาวิทยาลัย)25ธ.ค.2566"
            },
            "final": {
              "date": "14",
              "month": "Feb",
              "times": "13:00-16:00",
              "year": "2567",
              "room": "อาคารBห้องN(สอบตามตารางมหาวิทยาลัย)14ก.พ.2567"
            }
          }
        }
      ]
    }
  ]
}
```

## Tech Stack

[Go Fiber](https://docs.gofiber.io/), [Colly](https://github.com/gocolly/colly), [Cache](https://github.com/patrickmn/go-cache)

<!-- [**Redis**](https://redis.io/) : cache data -->
