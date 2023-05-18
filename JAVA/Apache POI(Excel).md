# Apache POI (Excel)

업무중에 데이터를 엑셀로 넘겨야하는 업무가 생겼다. 이걸 손으로 다 만드는건 무리가 있으므로... POI를 사용하기로 했다.

```Apache POI```는 아파치 재단에서 만든 라이브러리로 마이크로소프트 오피스 파일을 자바로 읽고 씨는 기능을 제공한다.

## JDK별 호환성
참고: [https://poi.apache.org/devel/](https://poi.apache.org/devel/)

- POI 4.0 and later require JDK version 1.8 or later.
- POI 3.11 and later 3.x versions require JDK version 1.6 or later.
- POI 3.5 to 3.10 required the JDK version 1.5 or later. Versions prior to 3.5 required JDK 1.4+.

나는 1.8 버전의 JDK를 사용하므로 3.x 또는 4.x 사용하려고한다.(23.05.18 기준 5.2.3버전까지 나와있다.)


## 의존성 추가

### 1. poi

- HSSF
- Excel 97(~2007) 파일 포맷을 사용할 때
- 옛날 엑셀버전을 읽고 쓸 경우

### 2. poi-ooxml

- XSSF
- Excel 2007 OOXML(*.xlsx) 파일 포맷을 사용할 때
- 요즘 주로 사용되는 엑셀을 읽고 쓸 경우

### 3. poi/poi-ooxml 차이

|poi|poi-ooxml|
|---|---|
|HSSFWorkbook|XSSFWorkbook|
|HSSFSheet|XSSFSheet|

### 4. 의존성 추가

```xml
<!-- https://mvnrepository.com/artifact/org.apache.poi/poi -->
<dependency>
    <groupId>org.apache.poi</groupId>
    <artifactId>poi</artifactId>
    <version>4.1.2</version>
</dependency>

<!-- https://mvnrepository.com/artifact/org.apache.poi/poi-ooxml -->
<dependency>
    <groupId>org.apache.poi</groupId>
    <artifactId>poi-ooxml</artifactId>
    <version>4.1.2</version>
</dependency>
```
---

## Excel 파일 읽기

테스트 파일의 내용을 아래와 같다.(XSSF를 사용한 예제이다.)

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/5c707da8-97b5-4f83-941c-a9155c6b9cc6)

설명은 코드내 주석으로 달아두었다.

```java
package org.example;

import org.apache.poi.xssf.usermodel.XSSFCell;
import org.apache.poi.xssf.usermodel.XSSFRow;
import org.apache.poi.xssf.usermodel.XSSFSheet;
import org.apache.poi.xssf.usermodel.XSSFWorkbook;

import java.io.FileInputStream;

public class Main {
    public static void main(String[] args) {

        try{
            // 해당 경로에 있는 파일일을 읽는다.
            FileInputStream file = new FileInputStream("E:/workspace/excelTest/ExcelRead.xlsx");
            // 파일을 엑셀 형태로 바꾼다?? 표현을 모르겠음
            XSSFWorkbook workbook = new XSSFWorkbook(file);

            // 파일을 workbook형태로 바꾸었으니 필요없는 리소스는 close
            file.close();

            // 엑셀의 시트를 의미한다. 0번은 첫번째 시트를 의미한다.
            XSSFSheet sheet = workbook.getSheetAt(0);

            // 최대 Row를 의미한다.
            int maxRow = sheet.getPhysicalNumberOfRows();
            for(int rowNum = 0; rowNum < maxRow; rowNum++){

                // row에는 1개의 행이 들어간다.
                XSSFRow row = sheet.getRow(rowNum);

                // row데이터가 없는경우는 continue
                if(row == null) continue;

                // 최대 Cell의 개수가 들어간다.
                int maxCell = row.getPhysicalNumberOfCells();

                for(int cellNum = 0; cellNum < maxCell; cellNum++){
                    XSSFCell cell = row.getCell(cellNum);

                    // cell데이터가 없는경우 continue
                    if(cell == null) continue;

                    String cellValue = "";

                    // cell 마다 타입이 있어서 해당 타입에 맞게 읽는다.
                    switch (cell.getCellType()){
                        case FORMULA:
                            cellValue = cell.getCellFormula();
                            break;
                        case NUMERIC:
                            cellValue = cell.getNumericCellValue() + "";
                            break;
                        case STRING:
                            cellValue = cell.getStringCellValue() + "";
                            break;
                        case BLANK:
                            cellValue = cell.getBooleanCellValue() + "";
                            break;
                        case ERROR:
                            cellValue = cell.getErrorCellValue() + "";
                            break;
                    }
                    System.out.printf("%s / ", cellValue);
                }
                System.out.println();
            }

        }catch (Exception e){
            e.printStackTrace();
        }
    }
}
```

### 출력

출력은 그냥 보여주기 위함일뿐 따로 출력을 이쁘게 꾸미진 않았다.
        
```
이름 / 키 / 몸무게 / 
홍길동 / 167.0 / 56.0 / 
둘리 / 140.0 / 45.0 / 
```

---

## Excel 파일 만들기

```java
package org.example;

import org.apache.poi.xssf.usermodel.XSSFCell;
import org.apache.poi.xssf.usermodel.XSSFRow;
import org.apache.poi.xssf.usermodel.XSSFSheet;
import org.apache.poi.xssf.usermodel.XSSFWorkbook;

import java.io.File;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Main {
    public static void main(String[] args) {

        // 파일을 생성할 경로와 파일이름
        String filePath = "E:/workspace/excelTest/ExcelWrite.xlsx";

        // 엑셀에 넣을 데이터를 리스트형식으로 만들어 주었다.
        List<List<String>> inputData = new ArrayList<>();
        inputData.add(Arrays.asList("이름", "키", "몸무게"));
        inputData.add(Arrays.asList("백설공주", "162", "50"));
        inputData.add(Arrays.asList("유재석", "178", "75"));
        
        // workbook을 만들어준다.
        XSSFWorkbook workbook = new XSSFWorkbook();

        // workbook에서 sheet도 만들어준다.
        XSSFSheet sheet = workbook.createSheet("Test Sheet");

        // rowNum은 몇번째 줄을 만들것인지 명시하기 위해
        int rowNum = 0;
        for(List<String> list : inputData) {
            // rowNum번째 row을 생성한다.
            XSSFRow row = sheet.createRow(rowNum++);

            // cellNum은 열번호 A, B, C ...
            int cellNum = 0;
            for(String s : list){
                // cellNum번째 cell을 생성한다.
                XSSFCell cell = row.createCell(cellNum++);
                // 해당 셀에 데이터를 넣는다.
                cell.setCellValue(s);
            }
        }
        try{
            // 저장하는 부분!
            FileOutputStream out = new FileOutputStream(new File(filePath));
            workbook.write(out);
            out.close();
        }catch (Exception e){
            e.printStackTrace();
        }
    }
}
```

### 출력 파일 확인

![image](https://github.com/YoonSeok-Heo/TIL/assets/113662725/5b6901db-7e2d-4daa-8b7d-a6f116f20721)

아주 만족한다.~~


---


