const fs = require('fs');
const p = require('path');

const dartFiles = [];

function main() {
  const openapiDirPath = p.resolve(__dirname, '../../gateway_admin_sdk');
  searchDartFiles(openapiDirPath);
  for (const filePath of dartFiles) {
    fixFile(filePath);
    console.log('Fixed file:', filePath);
  }
}

function searchDartFiles(path) {
  const isDir = fs.lstatSync(path).isDirectory();
  if (isDir) {
    const dirContent = fs.readdirSync(path);
    for (const dirContentPath of dirContent) {
      const fullPath = p.resolve(path, dirContentPath);
      searchDartFiles(fullPath);
    }
  } else {
    if (path.includes('.dart')) {
      dartFiles.push(path);
    }
  }
}

function fixFile(path) {
  const fileContent = fs.readFileSync(path).toString();
  const fixedContent = fixOthers(fileContent);
  fs.writeFileSync(path, fixedContent);
}

const fixOthers = fileContent => {
  let content = fileContent;
  for (const entry of otherFixers.entries()) {
    content = content.replace(entry[0], entry[1]);
  }
  return content;
};

const otherFixers = new Map([
  // ? Base fixers for Dio and standard params
  [
    '// @dart=2.7',
    '// ',
  ],
  [
    /response\.request/gm,
    'response.requestOptions',
  ],
  [
    /request: /gm,
    'requestOptions: ',
  ],
  [
    /Iterable<Object> serialized/gm,
    'Iterable<Object?> serialized',
  ],
  [
    /(?<type>^ +Uint8List)(?<value> file,)/gm,
    '$<type>?$<value>',
  ],
  [
    /(?<type>^ +String)(?<value> additionalMetadata,)/gm,
    '$<type>?$<value>',
  ],
  [
    /(?<type>^ +ProgressCallback)(?<value> onReceiveProgress,)/gm,
    '$<type>?$<value>',
  ],
  [
    /(?<type>^ +ProgressCallback)(?<value> onSendProgress,)/gm,
    '$<type>?$<value>',
  ],
  [
    /(?<type>^ +ValidateStatus)(?<value> validateStatus,)/gm,
    '$<type>?$<value>',
  ],
  [
    /(?<type>^ +Map<String, dynamic>)(?<value> extra,)/gm,
    '$<type>?$<value>',
  ],
  [
    /(?<type>^ +Map<String, dynamic>)(?<value> headers,)/gm,
    '$<type>?$<value>',
  ],
  [
    /(?<type>^ +CancelToken)(?<value> cancelToken,)/gm,
    '$<type>?$<value>',
  ],
  [
    /(@nullable\n)(?<annotation>^ +@.*\n)(?<type>.*)(?<getter> get )(?<variable>.*\n)/gm,
    '$<annotation>$<spaces>$<type>?$<getter>$<variable>',
  ],
  [
    'final result = <Object>[];',
    'final result = <Object?>[];',
  ],
  [
    'Iterable<Object> serialize',
    'Iterable<Object?> serialize',
  ],
  [
    /^ *final _response = await _dio.request<dynamic>\(\n +_request\.path,\n +data: _bodyData,\n +options: _request,\n +\);/gm,
    `_request.data = _bodyData;
    
    final _response = await _dio.fetch<dynamic>(_request);
    `,
  ],
  // ? Special, custom params for concrete API
  [
    /(?<type>^ +String)(?<value> apiKey,)/gm,
    '$<type>?$<value>',
  ],
  [
    /(?<type>^ +String)(?<value> name,)/gm,
    '$<type>?$<value>',
  ],
  [
    /(?<type>^ +String)(?<value> status,)/gm,
    '$<type>?$<value>',
  ],
]);

main();