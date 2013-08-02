
void setup() {
  pinMode(3, OUTPUT);
  Serial.begin(9600);
}
double val = 0.0;
void loop() {
  val += 0.001;
  int sine = int((sin(val*5) + 1.0) * 256.0);
  int cosine = int((cos(val * 2.0) + 1.0) * 400.0) + 112;
  Serial.print(sine);
  Serial.print(":");
  Serial.print(cosine);
  Serial.print(":");
  Serial.print(analogRead(A0));
  Serial.print("\n");
  delay(2);
}
