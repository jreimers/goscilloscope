
void setup() {
  pinMode(3, OUTPUT);
  Serial.begin(9600);
}
double val = 0.0;
void loop() {
  val += 0.02;
  int sine = int((sin(val) + 1.0) * 512.0);
  int cosine = int((cos(val) + 1.0) * 512.0);
  Serial.print(sine);
  Serial.print(":");
  Serial.print(cosine);
  Serial.print("\n");
  delay(2);
}
