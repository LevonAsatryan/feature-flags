import { IsNotEmpty, IsString, MinLength } from 'class-validator';

export class CreateCompanyDto {
  @IsString()
  @MinLength(3)
  @IsNotEmpty()
  name: string;
}
